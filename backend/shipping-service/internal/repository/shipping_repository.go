package repository

import (
	"context"
	"errors"

	"shipping-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShippingRepository interface {
	GetShipments(ctx context.Context, ids []uuid.UUID) ([]model.Shipment, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Shipment, error)
	GetByOrderID(ctx context.Context, orderID uuid.UUID) (*model.Shipment, error)
	Create(ctx context.Context, shipment *model.Shipment) error
	Update(ctx context.Context, shipment *model.Shipment) error
}

type shippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) ShippingRepository {
	return &shippingRepository{db: db}
}

func (r *shippingRepository) GetShipments(
	ctx context.Context,
	ids []uuid.UUID,
) ([]model.Shipment, error) {
	var shipments []model.Shipment

	err := r.db.
		WithContext(ctx).
		Where("id in (?)", ids).
		Find(&shipments).
		Error

	return shipments, err
}

func (r *shippingRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Shipment, error) {
	var shipment model.Shipment

	err := r.db.
		WithContext(ctx).
		First(&shipment, "id = ?", id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &shipment, err
}

func (r *shippingRepository) GetByOrderID(
	ctx context.Context,
	orderID uuid.UUID,
) (*model.Shipment, error) {
	var shipment model.Shipment

	err := r.db.
		WithContext(ctx).
		Where("order_id = ?", orderID).
		First(&shipment).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &shipment, err
}

func (r *shippingRepository) Create(
	ctx context.Context,
	shipment *model.Shipment,
) error {
	return r.db.
		WithContext(ctx).
		Create(shipment).
		Error
}

func (r *shippingRepository) Update(
	ctx context.Context,
	shipment *model.Shipment,
) error {
	return r.db.
		WithContext(ctx).
		Model(&model.Shipment{}).
		Where("id = ?", shipment.ID).
		Updates(shipment).
		Error
}
