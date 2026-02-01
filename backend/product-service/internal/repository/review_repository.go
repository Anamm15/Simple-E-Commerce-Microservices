package repository

import (
	"context"

	"product-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetAllReviews(ctx context.Context) ([]model.Review, error)
	GetReviewByProductId(ctx context.Context, productId uuid.UUID) ([]model.Review, error)
	CreateReview(ctx context.Context, review *model.Review) error
	UpdateReview(ctx context.Context, review *model.Review) error
	DeleteReview(ctx context.Context, reviewId uuid.UUID, userID uuid.UUID) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetAllReviews(ctx context.Context) ([]model.Review, error) {
	var reviews []model.Review
	if err := r.db.WithContext(ctx).
		Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *reviewRepository) GetReviewByProductId(ctx context.Context, productId uuid.UUID) ([]model.Review, error) {
	var reviews []model.Review
	if err := r.db.WithContext(ctx).
		Where("product_id = ?", productId).
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewRepository) CreateReview(ctx context.Context, review *model.Review) error {
	if err := r.db.WithContext(ctx).
		Create(&review).Error; err != nil {
		return err
	}
	return nil
}

func (r *reviewRepository) UpdateReview(ctx context.Context, review *model.Review) error {
	if err := r.db.WithContext(ctx).
		Save(&review).Error; err != nil {
		return err
	}

	return nil
}

func (r *reviewRepository) DeleteReview(ctx context.Context, reviewId uuid.UUID, userID uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Where("id = ? and user_id = ?", reviewId, userID).
		Delete(&model.Review{}).Error; err != nil {
		return err
	}
	return nil
}
