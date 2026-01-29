package service

import (
	"context"

	"product-service/internal/dto"
	"product-service/internal/helper/mapper"
	productpb "product-service/internal/pb/product"
	"product-service/internal/repository"
	"product-service/internal/util"

	"github.com/google/uuid"
)

type ProductService interface {
	GetProducts(ctx context.Context, metadata *dto.SearchProductRequestDTO) (productpb.ProductList, error)
	GetProductDetail(ctx context.Context, productID uuid.UUID) (productpb.ProductDetail, error)
	GetProductBatch(ctx context.Context, productIDs []uuid.UUID) (productpb.ProductBatchResponse, error)
	CreateProduct(ctx context.Context, request dto.CreateProductRequestDTO) (productpb.ProductDetail, error)
	UpdateProduct(ctx context.Context, request dto.UpdateProductRequestDTO) (productpb.ProductDetail, error)
	DeleteProduct(ctx context.Context, productID string) error
}

type productService struct {
	productRepository  repository.ProductRepository
	reviewRepository   repository.ReviewRepository
	categoryRepository repository.CategoryRepository
}

func NewProductService(
	productRepository repository.ProductRepository,
	reviewRepository repository.ReviewRepository,
	categoryRepository repository.CategoryRepository,
) ProductService {
	return &productService{
		productRepository:  productRepository,
		reviewRepository:   reviewRepository,
		categoryRepository: categoryRepository,
	}
}

func (r *productService) GetProducts(ctx context.Context, metadata *dto.SearchProductRequestDTO) (productpb.ProductList, error) {
	return productpb.ProductList{}, nil
}

func (r *productService) GetProductDetail(ctx context.Context, productID string) (productpb.ProductDetail, error) {
	productIDParsed, err := util.StringToUUID(productID)
	if err != nil {
		return productpb.ProductDetail{}, err
	}

	product, err := r.productRepository.GetDetailProduct(ctx, productIDParsed)
	if err != nil {
		return productpb.ProductDetail{}, err
	}

	return *mapper.MapToProductDetail(&product), nil
}

func (r *productService) GetProductBatch(ctx context.Context, productIDs []uuid.UUID) (productpb.ProductBatchResponse, error) {
}

func (r *productService) CreateProduct(ctx context.Context, request dto.CreateProductRequestDTO) (productpb.ProductDetail, error) {
}

func (r *productService) UpdateProduct(ctx context.Context, request dto.UpdateProductRequestDTO) (productpb.ProductDetail, error) {
}

func (r *productService) DeleteProduct(ctx context.Context, productID string) error {
	productIDParsed, err := util.StringToUUID(productID)
	if err != nil {
		return err
	}

	err = r.productRepository.DeleteProduct(ctx, productIDParsed)
	if err != nil {
		return err
	}

	return nil
}
