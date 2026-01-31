package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"inventory-service/internal/dto"
	"inventory-service/internal/helper/enum"
	"inventory-service/internal/helper/mapper"
	"inventory-service/internal/model"
	inventorypb "inventory-service/internal/pb/inventory"
	"inventory-service/internal/repository"
	"inventory-service/internal/util"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryService interface {
	CreateStock(ctx context.Context, request dto.CreateStockRequestDTO) (*inventorypb.StockCount, error)
	CheckStock(ctx context.Context, productID string) (*inventorypb.StockCount, error)
	ReserveStock(ctx context.Context, request dto.ReserveStockRequestDTO) (*inventorypb.ReserveStockResponse, error)
	UpdateStock(ctx context.Context, request dto.UpdateStockRequestDTO) (*inventorypb.StockCount, error)
	ReleaseStock(ctx context.Context, request dto.ReleaseStockRequestDTO) error
}

type inventoryService struct {
	db                       *gorm.DB
	inventoryRepo            repository.InventoryRepository
	inventoryReservationRepo repository.InventoryReservationRepository
}

func NewInventoryService(
	db *gorm.DB,
	inventoryRepo repository.InventoryRepository,
	inventoryReservationRepo repository.InventoryReservationRepository,
) InventoryService {
	return &inventoryService{
		db:                       db,
		inventoryRepo:            inventoryRepo,
		inventoryReservationRepo: inventoryReservationRepo,
	}
}

func (s *inventoryService) CreateStock(ctx context.Context, request dto.CreateStockRequestDTO) (*inventorypb.StockCount, error) {
	inventory := request.ToModel()
	err := s.inventoryRepo.Create(ctx, inventory)
	if err != nil {
		return nil, err
	}

	return mapper.MapToInventoryStockCount(inventory), nil
}

func (s *inventoryService) CheckStock(ctx context.Context, productID string) (*inventorypb.StockCount, error) {
	productIDParsed, err := util.StringToUUID(productID)
	if err != nil {
		return nil, err
	}

	inventory, err := s.inventoryRepo.GetByProductID(ctx, productIDParsed)
	if err != nil {
		return nil, err
	}

	return mapper.MapToInventoryStockCount(inventory), nil
}

func (s *inventoryService) ReserveStock(ctx context.Context, request dto.ReserveStockRequestDTO) (*inventorypb.ReserveStockResponse, error) {
	productIDs := make([]uuid.UUID, 0, len(request.Items))
	itemMap := make(map[uuid.UUID]int32) // Map ProductID -> Quantity

	for _, item := range request.Items {
		pid, err := util.StringToUUID(item.ProductId)
		if err != nil {
			return nil, err
		}
		productIDs = append(productIDs, pid)
		itemMap[pid] = item.Quantity
	}

	orderIDParsed, err := util.StringToUUID(request.OrderID)
	if err != nil {
		return nil, err
	}

	// 1. Mulai Transaksi
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	// Defer rollback agar aman jika function panic atau return error di tengah
	defer tx.Rollback()

	// 2. Locking Read (SELECT ... FOR UPDATE)
	inventories, err := s.inventoryRepo.GetBatchInventoryByProductIDsWithLock(ctx, tx, productIDs)
	if err != nil {
		return nil, err
	}

	var outOfStockProductIDs []string
	var inventoryReservations []*model.InventoryReservation

	// 3. Map existing inventory untuk akses cepat
	dbInvMap := make(map[uuid.UUID]*model.Inventory)
	for _, inv := range inventories {
		dbInvMap[inv.ProductID] = inv
	}

	// 4. Logic Validasi & Update Memory
	for pid, requestedQty := range itemMap {
		inv, exists := dbInvMap[pid]

		if !exists || (inv.TotalStock-inv.ReservedStock) < requestedQty {
			outOfStockProductIDs = append(outOfStockProductIDs, pid.String())
			continue
		}

		inv.ReservedStock += requestedQty
		inventoryReservations = append(inventoryReservations, &model.InventoryReservation{
			InventoryID: inv.ID,
			ProductID:   inv.ProductID,
			OrderID:     orderIDParsed,
			Quantity:    requestedQty,
			ExpiryAt:    time.Now().Add(time.Hour * 24),
		})
	}

	// 5. Decision: All-or-Nothing
	if len(outOfStockProductIDs) > 0 {
		return &inventorypb.ReserveStockResponse{
			OutOfStockProductIds: outOfStockProductIDs,
		}, nil
	}

	// 6. Eksekusi Update ke DB
	// Update Inventory Master
	for _, inv := range dbInvMap {
		if _, exists := itemMap[inv.ProductID]; exists {
			if err := s.inventoryRepo.UpdateWithTx(ctx, tx, inv); err != nil {
				return nil, err
			}
		}
	}

	// Insert Reservation Log (Batch Insert)
	if err := s.inventoryReservationRepo.CreateBatchWithTx(ctx, tx, inventoryReservations); err != nil {
		return nil, err
	}

	// 7. Commit Transaksi
	if err := tx.Commit(); err != nil {
		return nil, err.Error
	}

	return &inventorypb.ReserveStockResponse{}, nil
}

func (s *inventoryService) UpdateStock(ctx context.Context, request dto.UpdateStockRequestDTO) (*inventorypb.StockCount, error) {
	productID, err := util.StringToUUID(request.ProductID)
	if err != nil {
		return &inventorypb.StockCount{}, err
	}

	inventory, err := s.inventoryRepo.GetByProductID(ctx, productID)
	if err != nil {
		return &inventorypb.StockCount{}, err
	}

	if request.Quantity < 0 {
		return &inventorypb.StockCount{}, nil
	}

	inventory.TotalStock = request.Quantity

	err = s.inventoryRepo.Update(ctx, inventory)
	if err != nil {
		return &inventorypb.StockCount{}, err
	}

	return mapper.MapToInventoryStockCount(inventory), nil
}

func (s *inventoryService) ReleaseStock(ctx context.Context, request dto.ReleaseStockRequestDTO) error {
	orderID, err := util.StringToUUID(request.OrderID)
	if err != nil {
		return err
	}

	// 1. Mulai Transaksi
	// FIX: Begin() mengembalikan *gorm.DB, error harus dicek dari object db atau return value tergantung driver,
	// tapi di GORM v2 biasanya error akses via .Error
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	// 2. Tentukan Status Target (Status AKHIR)
	var finalStatus enum.InventoryReservationStatus
	if request.Action == "CANCEL" {
		finalStatus = enum.InventoryReservationStatusReleased
	} else if request.Action == "COMMIT" {
		finalStatus = enum.InventoryReservationStatusCommited
	} else {
		return errors.New("invalid action")
	}

	// 3. Cari Reservasi yang masih PENDING/RESERVED
	currentStatus := enum.InventoryReservationStatusReserved
	reservations, err := s.inventoryReservationRepo.GetByOrderID(ctx, tx, orderID, &currentStatus)
	if err != nil {
		return err
	}

	if len(reservations) == 0 {
		// Idempotency: Mungkin sudah diproses sebelumnya? Return nil atau error tergantung bisnis.
		return nil
	}

	// Kumpulkan Product ID
	var productIDs []uuid.UUID
	for _, ir := range reservations {
		productIDs = append(productIDs, ir.ProductID)
	}

	// 4. Lock Inventory Master
	inventories, err := s.inventoryRepo.GetBatchInventoryByProductIDsWithLock(ctx, tx, productIDs)
	if err != nil {
		return err
	}

	// Helper Map untuk akses O(1) -> Menghindari Nested Loop
	invMap := make(map[uuid.UUID]*model.Inventory)
	for _, inv := range inventories {
		invMap[inv.ProductID] = inv
	}

	// 5. Logic Perhitungan Stok (FIXED)
	for _, ir := range reservations {
		inv, exists := invMap[ir.ProductID]
		if !exists {
			return fmt.Errorf("inventory not found for product %s", ir.ProductID)
		}

		if request.Action == "CANCEL" {
			// BATAL: Kembalikan jatah reserved ke pool bebas.
			// Total Stock TETAP (karena barang belum keluar gudang).
			inv.ReservedStock -= ir.Quantity

			// Safety check agar tidak minus
			if inv.ReservedStock < 0 {
				inv.ReservedStock = 0
			}

		} else if request.Action == "COMMIT" {
			// TERJUAL: Barang keluar gudang.
			// Kurangi Total Stock (Fisik berkurang).
			inv.TotalStock -= ir.Quantity
			// Kurangi Reserved Stock (Kewajiban booking hilang).
			inv.ReservedStock -= ir.Quantity

			// Safety check
			if inv.TotalStock < 0 {
				return errors.New("negative stock detected")
			}
		}

		// Update status di object reservation struct (untuk update step 7)
		ir.Status = finalStatus
	}

	// 6. Update Inventory Master
	for _, inv := range invMap {
		if err := s.inventoryRepo.UpdateWithTx(ctx, tx, inv); err != nil {
			return err
		}
	}

	// 7. Update Inventory Reservation Status
	for _, ir := range reservations {
		if err := s.inventoryReservationRepo.UpdateStatusWithTx(ctx, tx, ir); err != nil {
			return err
		}
	}

	// 8. Commit
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
