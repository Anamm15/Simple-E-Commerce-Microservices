package repository

import (
	"context"
	"errors"

	"cart-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetCart(ctx context.Context, userID string) ([]model.CartItem, error)
	AddItem(ctx context.Context, cartItem *model.CartItem) error
	RemoveItem(ctx context.Context, userID uuid.UUID, productID uuid.UUID) error
	UpdateItem(ctx context.Context, userId uuid.UUID, productId uuid.UUID, quantity int32) error
	ClearCart(ctx context.Context, userID string) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetCart(ctx context.Context, userID string) ([]model.CartItem, error) {
	var item []model.CartItem

	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at ASC").
		Find(&item).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []model.CartItem{}, nil
	}

	return item, err
}

func (r *cartRepository) AddItem(ctx context.Context, cartItem *model.CartItem) error {
	return r.db.
		WithContext(ctx).
		Create(cartItem).
		Error
}

func (r *cartRepository) RemoveItem(
	ctx context.Context,
	userID uuid.UUID,
	productID uuid.UUID,
) error {
	result := r.db.
		WithContext(ctx).
		Delete(
			&model.CartItem{},
			"user_id = ? AND product_id = ?",
			userID,
			productID,
		)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (r *cartRepository) UpdateItem(
	ctx context.Context,
	userId uuid.UUID,
	productId uuid.UUID,
	quantity int32,
) error {
	return r.db.
		WithContext(ctx).
		Model(&model.CartItem{}).
		Where("user_id = ? AND product_id = ?", userId, productId).
		Update("quantity", quantity).
		Error
}

func (r *cartRepository) ClearCart(ctx context.Context, userID string) error {
	return r.db.
		WithContext(ctx).
		Delete(&model.CartItem{}, "user_id = ?", userID).
		Error
}
