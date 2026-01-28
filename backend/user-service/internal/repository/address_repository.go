package repository

import (
	"context"
	"errors"

	"user-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Address, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]model.Address, error)
	Create(ctx context.Context, address *model.Address) error
	Update(ctx context.Context, address *model.Address) error
	Delete(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		db: db,
	}
}

func (r *addressRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Address, error) {
	var address model.Address

	err := r.db.
		WithContext(ctx).
		First(&address, "id = ?", id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &address, err
}

func (r *addressRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]model.Address, error) {
	var addresses []model.Address

	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Order("is_primary DESC, created_at ASC").
		Find(&addresses).
		Error

	return addresses, err
}

func (r *addressRepository) Create(ctx context.Context, address *model.Address) error {
	return r.db.
		WithContext(ctx).
		Create(address).
		Error
}

func (r *addressRepository) Update(ctx context.Context, address *model.Address) error {
	return r.db.
		WithContext(ctx).
		Model(&model.Address{}).
		Where("id = ?", address.ID).
		Updates(address).
		Error
}

func (r *addressRepository) Delete(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Delete(&model.Address{}, "id = ? AND user_id = ?", id, userID)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
