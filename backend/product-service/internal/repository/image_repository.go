package repository

import (
	"context"

	"product-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImageRepository interface {
	GetImageByProductID(ctx context.Context, productID uuid.UUID) (model.Image, error)
	CreateImage(ctx context.Context, image *model.Image) error
	CreateBatchImage(ctx context.Context, images []model.Image) error
	DeleteImage(ctx context.Context, imageID uuid.UUID) (model.Image, error)
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{db: db}
}

func (r *imageRepository) GetImageByProductID(ctx context.Context, productID uuid.UUID) (model.Image, error) {
	var image model.Image
	if err := r.db.WithContext(ctx).
		Where("product_id = ?", productID).
		Find(&image).Error; err != nil {
		return model.Image{}, err
	}
	return image, nil
}

func (r *imageRepository) CreateImage(ctx context.Context, image *model.Image) error {
	if err := r.db.WithContext(ctx).
		Create(&image).Error; err != nil {
		return err
	}
	return nil
}

func (r *imageRepository) CreateBatchImage(ctx context.Context, images []model.Image) error {
	if err := r.db.WithContext(ctx).
		Create(&images).Error; err != nil {
		return err
	}
	return nil
}

func (r *imageRepository) DeleteImage(ctx context.Context, imageID uuid.UUID) (model.Image, error) {
	var image model.Image
	if err := r.db.WithContext(ctx).
		Where("id = ?", imageID).
		Delete(&image).Error; err != nil {
		return model.Image{}, err
	}
	return image, nil
}
