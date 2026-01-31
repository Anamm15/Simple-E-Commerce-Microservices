package controller

import (
	"context"

	"product-service/internal/dto"
	"product-service/internal/helper/mapper"
	productpb "product-service/internal/pb/product"
	"product-service/internal/service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type productController struct {
	productpb.UnimplementedProductServiceServer
	productService service.ProductService
	imageService   service.ImageService
}

func NewProductController(productService service.ProductService, imageService service.ImageService) productpb.ProductServiceServer {
	return &productController{productService: productService, imageService: imageService}
}

func (r *productController) GetProducts(ctx context.Context, request *productpb.SearchFilterRequest) (*productpb.ProductList, error) {
	input := dto.SearchProductRequestDTO{
		SearchQuery:  request.SearchQuery,
		CategorySlug: request.CategorySlug,
		MinPrice:     request.MinPrice,
		MaxPrice:     request.MaxPrice,
		Page:         request.Page,
		Limit:        request.Limit,
	}

	return r.productService.GetProducts(ctx, input)
}

func (r *productController) GetProductDetail(ctx context.Context, request *productpb.GetProductDetailRequest) (*productpb.ProductDetail, error) {
	return r.productService.GetProductDetail(ctx, request.Id)
}

func (r *productController) GetBatchProduct(ctx context.Context, request *productpb.GetProductBatchRequest) (*productpb.ProductBatchResponse, error) {
	return r.productService.GetProductBatch(ctx, request.ProductIds)
}

func (r *productController) CreateProduct(ctx context.Context, request *productpb.CreateProductRequest) (*productpb.ProductDetail, error) {
	input := dto.CreateProductRequestDTO{
		Name:           request.Name,
		Description:    request.Description,
		Price:          request.Price,
		WeightG:        request.WeightG,
		Categories:     request.Categories,
		Thumbnail:      *mapper.MapImagePBToDTO(request.Thumbnail),
		AdditionalImgs: mapper.MapImagesListToDTO(request.AdditionalImgs),
		InitialStock:   request.InitialStock,
	}

	return r.productService.CreateProduct(ctx, input)
}

func (r *productController) UpdateProduct(ctx context.Context, request *productpb.UpdateProductRequest) (*productpb.ProductDetail, error) {
	input := dto.UpdateProductRequestDTO{
		Name:        &request.Name,
		Description: &request.Description,
		Price:       &request.Price,
		WeightG:     &request.WeightG,
		ProductID:   request.Id,
	}

	return r.productService.UpdateProduct(ctx, input)
}

func (r *productController) AddImageProduct(ctx context.Context, request *productpb.AddImageProductRequest) (*productpb.ImageProductResponse, error) {
	input := dto.AddImageProductRequestDTO{
		ProductID: request.ProductId,
		Images:    mapper.MapImagesListToDTO(request.Images),
	}

	return r.imageService.AddImageProduct(ctx, input)
}

func (r *productController) UpdateThumbnailProduct(ctx context.Context, request *productpb.UpdateThumbnailProductRequest) (*productpb.UpdateThumbnailProductResponse, error) {
	input := dto.UpdateThumbnailProductRequest{
		ProductID: request.ProductId,
		Image:     *mapper.MapImagePBToDTO(request.Image),
	}

	return r.productService.UpdateThumbnailProduct(ctx, input)
}

func (r *productController) DeleteImageProduct(ctx context.Context, request *productpb.DeleteImageProductRequest) (*emptypb.Empty, error) {
	err := r.imageService.DeleteImageProduct(ctx, request.ImageId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (r *productController) DeleteProduct(ctx context.Context, request *productpb.DeleteProductRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, r.productService.DeleteProduct(ctx, request.Id)
}
