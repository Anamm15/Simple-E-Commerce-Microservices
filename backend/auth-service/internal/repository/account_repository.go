package repository

import (
	"context"
	"errors"

	"auth-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(ctx context.Context, account *model.Account) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Account, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) (*model.Account, error)
	FindByEmail(ctx context.Context, email string) (*model.Account, error)
	FindByUsername(ctx context.Context, username string) (*model.Account, error)
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(ctx context.Context, account *model.Account) error {
	return r.db.
		WithContext(ctx).
		Create(account).
		Error
}

func (r *accountRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	var account model.Account

	err := r.db.
		WithContext(ctx).
		First(&account, id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &account, err
}

func (r *accountRepository) FindByEmail(ctx context.Context, email string) (*model.Account, error) {
	var account model.Account

	err := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&account).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &account, err
}

func (r *accountRepository) FindByUsername(ctx context.Context, username string) (*model.Account, error) {
	var account model.Account

	err := r.db.
		WithContext(ctx).
		Where("username = ?", username).
		First(&account).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &account, err
}

func (r *accountRepository) FindByUserID(ctx context.Context, userID uuid.UUID) (*model.Account, error) {
	var account model.Account

	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		First(&account).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &account, err
}

func (r *accountRepository) Update(ctx context.Context, account *model.Account) error {
	return r.db.
		WithContext(ctx).
		Model(&model.Account{}).
		Where("id = ?", account.ID).
		Updates(account).
		Error
}

func (r *accountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.
		WithContext(ctx).
		Delete(&model.Account{}, id).
		Error
}
