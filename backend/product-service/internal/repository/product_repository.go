package repository

import (
	"context"

	"product-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	GetProductByID(ctx context.Context, productId uuid.UUID) (*model.Product, error)
	GetBatchProductByIDS(ctx context.Context, productIDs []uuid.UUID) ([]model.Product, error)
	GetProductsByCategory(ctx context.Context, categoryId uuid.UUID) ([]model.Product, error)
	GetDetailProduct(ctx context.Context, productId uuid.UUID) (*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, product *model.Product) error
	UpdateThumnailProduct(ctx context.Context, productID uuid.UUID, thumbnail string) error
	DeleteProduct(ctx context.Context, productId uuid.UUID) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product

	if err := r.db.WithContext(ctx).
		Preload("Categories").
		Preload("Images").
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetProductByID(ctx context.Context, productId uuid.UUID) (*model.Product, error) {
	var product model.Product
	if err := r.db.WithContext(ctx).
		Where("id = ?", productId).
		Find(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetBatchProductByIDS(ctx context.Context, productIDs []uuid.UUID) ([]model.Product, error) {
	if len(productIDs) == 0 {
		return []model.Product{}, nil
	}

	var products []model.Product

	err := r.db.WithContext(ctx).
		Where("id IN ?", productIDs).
		Find(&products).
		Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetDetailProduct(ctx context.Context, productId uuid.UUID) (*model.Product, error) {
	var product model.Product
	if err := r.db.WithContext(ctx).
		Preload("Categories").
		Preload("Images").
		Where("id = ?", productId).
		Find(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetProductsByCategory(ctx context.Context, categoryId uuid.UUID) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.WithContext(ctx).
		Preload("Categories").
		Preload("Images").
		Where("category_id = ?", categoryId).
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) CreateProduct(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).
		Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).
		Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) UpdateThumnailProduct(ctx context.Context, productID uuid.UUID, thumbnail string) error {
	if err := r.db.WithContext(ctx).
		Where("id = ?", productID).
		Update("thumbnail", thumbnail).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, productId uuid.UUID) error {
	if err := r.db.WithContext(ctx).
		Where("id = ?", productId).
		Delete(&model.Product{}).Error; err != nil {
		return err
	}
	return nil
}
