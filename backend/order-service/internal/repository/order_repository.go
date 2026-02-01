package repository

import (
	"context"

	"order-service/internal/helper"
	"order-service/internal/helper/enum"
	"order-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAll(ctx context.Context, limit *int32, offset *int32, sort *string, filter *enum.OrderStatus) ([]model.Order, int64, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit *int32, offset *int32, sort *string, filter *enum.OrderStatus) ([]model.Order, int64, error)
	GetByID(ctx context.Context, orderID uuid.UUID) (*model.Order, error)
	GetDetailOrder(ctx context.Context, orderID uuid.UUID) (*model.Order, error)
	Create(ctx context.Context, order *model.Order) error
	Update(ctx context.Context, order *model.Order) error
	Delete(ctx context.Context, orderID uuid.UUID) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAll(
	ctx context.Context,
	limit *int32,
	offset *int32,
	sort *string,
	filter *enum.OrderStatus,
) ([]model.Order, int64, error) {
	var orders []model.Order
	var totalCount int64

	db := r.db.WithContext(ctx).
		Model(&model.Order{})

	db = helper.ApplyQueryOptions(db, limit, offset, sort, filter)

	if err := db.
		Count(&totalCount).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, totalCount, nil
}

func (r *orderRepository) GetByUserID(
	ctx context.Context,
	userID uuid.UUID,
	limit *int32,
	offset *int32,
	sort *string,
	filter *enum.OrderStatus,
) ([]model.Order, int64, error) {
	var orders []model.Order
	var totalCount int64

	db := r.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("user_id = ?", userID)

	db = helper.ApplyQueryOptions(db, limit, offset, sort, filter)

	if err := db.
		Count(&totalCount).
		Error; err != nil {
		return nil, 0, err
	}

	if err := db.Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, totalCount, nil
}

func (r *orderRepository) GetByID(
	ctx context.Context,
	orderID uuid.UUID,
) (*model.Order, error) {
	var order model.Order

	if err := r.db.WithContext(ctx).
		First(&order, "id = ?", orderID).
		Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepository) GetDetailOrder(
	ctx context.Context,
	orderID uuid.UUID,
) (*model.Order, error) {
	var order model.Order

	if err := r.db.WithContext(ctx).
		Preload("Items").
		First(&order, "id = ?", orderID).
		Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepository) Create(
	ctx context.Context,
	order *model.Order,
) error {
	return r.db.WithContext(ctx).
		Create(order).
		Error
}

func (r *orderRepository) Update(
	ctx context.Context,
	order *model.Order,
) error {
	return r.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("id = ?", order.ID).
		Updates(order).
		Error
}

func (r *orderRepository) Delete(
	ctx context.Context,
	orderID uuid.UUID,
) error {
	return r.db.WithContext(ctx).
		Delete(&model.Order{}, "id = ?", orderID).
		Error
}
