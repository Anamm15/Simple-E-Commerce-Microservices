package repository

import (
	"context"

	"payment-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	FindByOrderID(ctx context.Context, orderID uuid.UUID) (*model.Payment, error)
	Create(ctx context.Context, payment *model.Payment) error
	Update(ctx context.Context, payment *model.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) FindByOrderID(ctx context.Context, orderID uuid.UUID) (*model.Payment, error) {
	var payment model.Payment
	if err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) Create(ctx context.Context, payment *model.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *paymentRepository) Update(ctx context.Context, payment *model.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}
