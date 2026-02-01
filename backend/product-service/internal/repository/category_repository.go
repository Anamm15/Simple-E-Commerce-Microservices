package repository

import (
	"context"
	"errors"

	"product-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) ([]model.Category, error)
	AddProductCategory(ctx context.Context, productID uuid.UUID, categoryIDs []uuid.UUID) ([]model.Category, error)
	CreateCategory(ctx context.Context, category *model.Category) error
	UpdateCategory(ctx context.Context, category *model.Category) error
	DeleteCategory(ctx context.Context, categoryId uuid.UUID) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetAllCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.WithContext(ctx).
		Model(&model.Category{}).
		Select("id", "name").
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) AddProductCategory(ctx context.Context, productID uuid.UUID, categoryIDs []uuid.UUID) ([]model.Category, error) {
	var categories []model.Category

	if err := r.db.WithContext(ctx).
		Where("id IN ?", categoryIDs).
		Find(&categories).Error; err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, errors.New("category not found")
	}

	var product model.Product
	product.ID = productID

	if err := r.db.WithContext(ctx).
		Model(&product).
		Association("Category").
		Append(&categories); err != nil {
		return nil, err
	}

	var categoryResponseDTOs []model.Category
	for _, category := range categories {
		categoryResponseDTOs = append(categoryResponseDTOs, model.Category{ID: category.ID, Name: category.Name})
	}
	return categoryResponseDTOs, nil
}

func (r *categoryRepository) CreateCategory(ctx context.Context, category *model.Category) error {
	if err := r.db.WithContext(ctx).
		Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *model.Category) error {
	if err := r.db.WithContext(ctx).
		Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, categoryId uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Where("id = ?", categoryId).
		Delete(&model.Category{}).Error; err != nil {
		return err
	}
	return nil
}
