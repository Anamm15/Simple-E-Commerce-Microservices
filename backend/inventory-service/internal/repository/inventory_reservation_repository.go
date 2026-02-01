package repository

import (
	"context"

	"inventory-service/internal/helper/enum"
	"inventory-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InventoryReservationRepository interface {
	GetByOrderID(ctx context.Context, tx *gorm.DB, orderID uuid.UUID, status *enum.InventoryReservationStatus) ([]*model.InventoryReservation, error)
	Create(ctx context.Context, inventoryReservation *model.InventoryReservation) error
	CreateBatchWithTx(ctx context.Context, tx *gorm.DB, inventoryReservations []*model.InventoryReservation) error
	UpdateStatusWithTx(ctx context.Context, tx *gorm.DB, inventoryReservation *model.InventoryReservation) error
}

type inventoryReservationRepository struct {
	db *gorm.DB
}

func NewInventoryReservationRepository(db *gorm.DB) InventoryReservationRepository {
	return &inventoryReservationRepository{db: db}
}

func (r *inventoryReservationRepository) GetByOrderID(ctx context.Context, tx *gorm.DB, orderID uuid.UUID, status *enum.InventoryReservationStatus) ([]*model.InventoryReservation, error) {
	var inventoryReservations []*model.InventoryReservation

	query := tx.
		WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("order_id = ?", orderID)

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	err := query.Find(&inventoryReservations).Error
	return inventoryReservations, err
}

func (r *inventoryReservationRepository) Create(ctx context.Context, inventoryReservation *model.InventoryReservation) error {
	return r.db.WithContext(ctx).
		Create(inventoryReservation).Error
}

func (r *inventoryReservationRepository) CreateBatchWithTx(ctx context.Context, tx *gorm.DB, inventoryReservations []*model.InventoryReservation) error {
	return tx.WithContext(ctx).
		Create(inventoryReservations).Error
}

func (r *inventoryReservationRepository) UpdateStatusWithTx(ctx context.Context, tx *gorm.DB, inventoryReservation *model.InventoryReservation) error {
	return tx.WithContext(ctx).
		Model(&model.InventoryReservation{}).
		Where("id = ?", inventoryReservation.ID).
		Update("status", inventoryReservation.Status).Error
}
