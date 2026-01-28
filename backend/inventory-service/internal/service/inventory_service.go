package service

import (
	"context"

	"inventory-service/internal/dto"
	"inventory-service/internal/helper/mapper"
	inventorypb "inventory-service/internal/pb/inventory"
	"inventory-service/internal/repository"
	"inventory-service/internal/util"

	"github.com/google/uuid"
)

type InventoryService interface {
	CreateStock(ctx context.Context, request dto.CreateStockRequestDTO) (*inventorypb.StockCount, error)
	CheckStock(ctx context.Context, productID string) (*inventorypb.StockCount, error)
	ReserveStock(ctx context.Context, request dto.ReserveAndReleaseStockRequestDTO) (*inventorypb.ReserveStockResponse, error)
	UpdateStock(ctx context.Context, request dto.UpdateStockRequestDTO) (*inventorypb.StockCount, error)
	ReleaseStock(ctx context.Context, request dto.ReserveAndReleaseStockRequestDTO) error
}

type inventoryService struct {
	inventoryRepo repository.InventoryRepository
}

func NewInventoryService(inventoryRepo repository.InventoryRepository) InventoryService {
	return &inventoryService{inventoryRepo: inventoryRepo}
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

func (s *inventoryService) ReserveStock(ctx context.Context, request dto.ReserveAndReleaseStockRequestDTO) (*inventorypb.ReserveStockResponse, error) {
	var productIDs []uuid.UUID

	for _, item := range request.Items {
		productID, err := util.StringToUUID(item.ProductId)
		if err != nil {
			return nil, err
		}

		productIDs = append(productIDs, productID)
	}

	return nil, nil
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

	inventory.Stock = request.Quantity

	err = s.inventoryRepo.Update(ctx, inventory)
	if err != nil {
		return &inventorypb.StockCount{}, err
	}

	return mapper.MapToInventoryStockCount(inventory), nil
}

func (s *inventoryService) ReleaseStock(ctx context.Context, request dto.ReserveAndReleaseStockRequestDTO) error {
	var productIDs []uuid.UUID

	for _, item := range request.Items {
		productID, err := util.StringToUUID(item.ProductId)
		if err != nil {
			return err
		}

		productIDs = append(productIDs, productID)
	}

	return nil
}
