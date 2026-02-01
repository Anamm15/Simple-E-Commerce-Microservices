package repository

import (
	"context"
	"errors"

	"user-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	err := r.db.
		WithContext(ctx).
		Preload("Addresses").
		Find(&users).
		Error

	return users, err
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user model.User

	err := r.db.
		WithContext(ctx).
		Preload("Addresses").
		First(&user, "id = ?", id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.
		WithContext(ctx).
		Create(user).
		Error
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", user.ID).
		Updates(user).
		Error
}
