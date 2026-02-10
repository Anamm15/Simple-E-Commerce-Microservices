package repository

import (
	"context"

	"product-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context, limit int32, offset int32, search string, category uuid.UUID, minPrice int64, maxPrice int64, sort string) ([]model.Product, int64, error)
	GetProductByID(ctx context.Context, productId uuid.UUID) (*model.Product, error)
	GetBatchProductByIDS(ctx context.Context, productIDs []uuid.UUID) ([]model.Product, error)
	GetProductsByCategory(ctx context.Context, categoryId uuid.UUID) ([]model.Product, error)
	GetDetailProduct(ctx context.Context, productId uuid.UUID) (*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, product *model.Product) error
	UpdateThumnailProduct(ctx context.Context, productID uuid.UUID, thumbnail string, publicID string) error
	DeleteProduct(ctx context.Context, productId uuid.UUID) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts(
	ctx context.Context,
	limit int32,
	offset int32,
	search string,
	category uuid.UUID,
	minPrice int64,
	maxPrice int64,
	sort string,
) ([]model.Product, int64, error) {
	var (
		products []model.Product
		total    int64
	)

	baseQuery := r.db.WithContext(ctx).
		Model(&model.Product{})

	if search != "" {
		baseQuery = baseQuery.Where("products.name ILIKE ?", "%"+search+"%")
	}

	baseQuery = baseQuery.Where(
		"products.price BETWEEN ? AND ?",
		minPrice,
		maxPrice,
	)

	if category != uuid.Nil {
		baseQuery = baseQuery.
			Joins("JOIN product_categories pc ON pc.product_id = products.id").
			Where("pc.category_id = ?", category)
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	switch sort {
	case "price_asc":
		baseQuery = baseQuery.Order("products.price ASC")
	case "price_desc":
		baseQuery = baseQuery.Order("products.price DESC")
	case "name_asc":
		baseQuery = baseQuery.Order("products.name ASC")
	case "name_desc":
		baseQuery = baseQuery.Order("products.name DESC")
	default:
		baseQuery = baseQuery.Order("products.created_at DESC")
	}

	err := baseQuery.
		Preload("Images").
		Preload("Categories").
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
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

func (r *productRepository) UpdateThumnailProduct(ctx context.Context, productID uuid.UUID, thumbnail string, publicID string) error {
	if err := r.db.WithContext(ctx).
		Model(&model.Product{}).
		Where("id = ?", productID).
		Updates(map[string]interface{}{
			"thumbnail":           thumbnail,
			"thumbnail_public_id": publicID,
		}).Error; err != nil {
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
