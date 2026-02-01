package service

import (
	"context"

	"product-service/internal/dto"
	"product-service/internal/helper"
	"product-service/internal/helper/mapper"
	"product-service/internal/infrastructure/cloud"
	"product-service/internal/model"
	inventorypb "product-service/internal/pb/inventory"
	productpb "product-service/internal/pb/product"
	"product-service/internal/repository"
	"product-service/pkg/util"

	"github.com/google/uuid"
)

type ProductService interface {
	GetProducts(ctx context.Context, metadata dto.SearchProductRequestDTO) (*productpb.ProductList, error)
	GetProductDetail(ctx context.Context, productID string) (*productpb.ProductDetail, error)
	GetProductBatch(ctx context.Context, productIDs []string) (*productpb.ProductBatchResponse, error)
	CreateProduct(ctx context.Context, request dto.CreateProductRequestDTO) (*productpb.ProductDetail, error)
	UpdateProduct(ctx context.Context, request dto.UpdateProductRequestDTO) (*productpb.ProductDetail, error)
	UpdateThumbnailProduct(ctx context.Context, request dto.UpdateThumbnailProductRequest) (*productpb.UpdateThumbnailProductResponse, error)
	DeleteProduct(ctx context.Context, productID string) error
}

type productService struct {
	productRepository  repository.ProductRepository
	reviewRepository   repository.ReviewRepository
	categoryRepository repository.CategoryRepository
	inventoryClient    inventorypb.InventoryServiceClient
	cloudStorage       cloud.CloudinaryService
}

func NewProductService(
	productRepository repository.ProductRepository,
	reviewRepository repository.ReviewRepository,
	categoryRepository repository.CategoryRepository,
	inventoryClient inventorypb.InventoryServiceClient,
	cloudStorage cloud.CloudinaryService,
) ProductService {
	return &productService{
		productRepository:  productRepository,
		reviewRepository:   reviewRepository,
		categoryRepository: categoryRepository,
		inventoryClient:    inventoryClient,
		cloudStorage:       cloudStorage,
	}
}

func (r *productService) GetProducts(ctx context.Context, metadata dto.SearchProductRequestDTO) (*productpb.ProductList, error) {
	return &productpb.ProductList{}, nil
}

func (r *productService) GetProductDetail(ctx context.Context, productID string) (*productpb.ProductDetail, error) {
	productIDParsed, err := util.StringToUUID(productID)
	if err != nil {
		return nil, err
	}

	product, err := r.productRepository.GetDetailProduct(ctx, productIDParsed)
	if err != nil {
		return nil, err
	}

	return mapper.MapToProductDetail(product), nil
}

func (r *productService) GetProductBatch(ctx context.Context, productIDs []string) (*productpb.ProductBatchResponse, error) {
	var productIDsParsed []uuid.UUID
	for _, productID := range productIDs {
		productIDParsed, err := util.StringToUUID(productID)
		if err != nil {
			return nil, err
		}
		productIDsParsed = append(productIDsParsed, productIDParsed)
	}

	products, err := r.productRepository.GetBatchProductByIDS(ctx, productIDsParsed)
	if err != nil {
		return nil, err
	}

	productMap := make(map[string]*productpb.Product, len(products))
	for _, product := range products {
		productMap[product.ID.String()] = mapper.MapToProduct(&product)
	}

	return &productpb.ProductBatchResponse{
		Products: productMap,
	}, nil
}

func (r *productService) CreateProduct(ctx context.Context, request dto.CreateProductRequestDTO) (*productpb.ProductDetail, error) {
	// thumbnail
	uniqueFileName := helper.GenerateRandomFilename()
	fileReader := util.ByteToIOReader(request.Thumbnail.Data)

	fileUrl, err := r.cloudStorage.UploadFile(ctx, fileReader, uniqueFileName)
	if err != nil {
		return nil, err
	}

	// additional product image
	var additionalImages []model.Image
	for _, image := range request.AdditionalImgs {
		uniqueFileName := helper.GenerateRandomFilename()
		fileReader := util.ByteToIOReader(image.Data)

		fileUrl, err = r.cloudStorage.UploadFile(ctx, fileReader, uniqueFileName)
		if err != nil {
			return nil, err
		}
		additionalImages = append(additionalImages, model.Image{URL: fileUrl})
	}

	product := request.ToModel(fileUrl, additionalImages)
	err = r.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	// create inventory
	_, err = r.inventoryClient.CreateStock(ctx, &inventorypb.CreateStockRequest{
		ProductId: product.ID.String(),
		Quantity:  request.InitialStock,
	})
	if err != nil {
		return nil, err
	}

	return mapper.MapToProductDetail(product), nil
}

func (r *productService) UpdateThumbnailProduct(ctx context.Context, request dto.UpdateThumbnailProductRequest) (*productpb.UpdateThumbnailProductResponse, error) {
	productID, err := util.StringToUUID(request.ProductID)
	if err != nil {
		return nil, err
	}

	err = r.cloudStorage.DeleteFile(ctx, request.Image.Filename)
	if err != nil {
		return nil, err
	}

	uniqueFileName := helper.GenerateRandomFilename()
	fileReader := util.ByteToIOReader(request.Image.Data)

	fileUrl, err := r.cloudStorage.UploadFile(ctx, fileReader, uniqueFileName)
	if err != nil {
		return nil, err
	}

	err = r.productRepository.UpdateThumnailProduct(ctx, productID, fileUrl)
	if err != nil {
		return nil, err
	}
	return &productpb.UpdateThumbnailProductResponse{
		Image: fileUrl,
	}, nil
}

func (r *productService) UpdateProduct(ctx context.Context, request dto.UpdateProductRequestDTO) (*productpb.ProductDetail, error) {
	productID, err := util.StringToUUID(request.ProductID)
	if err != nil {
		return nil, err
	}

	product, err := r.productRepository.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	if request.Name != nil {
		product.Name = *request.Name
	}

	if request.Description != nil {
		product.Description = *request.Description
	}

	if request.Price != nil {
		product.Price = *request.Price
	}

	if request.WeightG != nil {
		product.WeightG = *request.WeightG
	}

	err = r.productRepository.UpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return mapper.MapToProductDetail(product), nil
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
