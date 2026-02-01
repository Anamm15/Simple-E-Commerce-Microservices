package repository

import (
	"context"
	"errors"

	"inventory-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InventoryRepository interface {
	GetByProductID(ctx context.Context, productID uuid.UUID) (*model.Inventory, error)
	GetBatchInventoryByProductIDsWithLock(ctx context.Context, tx *gorm.DB, productIDs []uuid.UUID) ([]*model.Inventory, error)
	Create(ctx context.Context, inventory *model.Inventory) error
	Update(ctx context.Context, inventory *model.Inventory) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, inventory *model.Inventory) error
	Delete(ctx context.Context, productID uuid.UUID) error
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetByProductID(
	ctx context.Context,
	productID uuid.UUID,
) (*model.Inventory, error) {
	var inventory model.Inventory

	err := r.db.
		WithContext(ctx).
		First(&inventory, "product_id = ?", productID).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &inventory, err
}

func (r *inventoryRepository) GetBatchInventoryByProductIDsWithLock(
	ctx context.Context,
	tx *gorm.DB,
	productIDs []uuid.UUID,
) ([]*model.Inventory, error) {
	var inventories []*model.Inventory

	err := tx.
		WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("product_id IN ?", productIDs).
		Find(&inventories).
		Error

	return inventories, err
}

func (r *inventoryRepository) Create(
	ctx context.Context,
	inventory *model.Inventory,
) error {
	return r.db.
		WithContext(ctx).
		Create(inventory).
		Error
}

func (r *inventoryRepository) Update(
	ctx context.Context,
	inventory *model.Inventory,
) error {
	return r.db.
		WithContext(ctx).
		Model(&model.Inventory{}).
		Where("product_id = ?", inventory.ProductID).
		Updates(map[string]interface{}{
			"stock": inventory.TotalStock,
		}).
		Error
}

func (r *inventoryRepository) UpdateWithTx(
	ctx context.Context,
	tx *gorm.DB,
	inventory *model.Inventory,
) error {
	return tx.WithContext(ctx).
		Model(&model.Inventory{}).
		Where("product_id = ?", inventory.ProductID).
		Updates(map[string]interface{}{
			"stock": inventory.TotalStock,
		}).
		Error
}

func (r *inventoryRepository) Delete(
	ctx context.Context,
	productID uuid.UUID,
) error {
	return r.db.
		WithContext(ctx).
		Delete(&model.Inventory{}, "product_id = ?", productID).
		Error
}
