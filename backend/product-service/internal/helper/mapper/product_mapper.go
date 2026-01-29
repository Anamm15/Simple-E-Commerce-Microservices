package mapper

import (
	"product-service/internal/model"
	productpb "product-service/internal/pb/product"
)

func MapToProduct(product *model.Product) *productpb.Product {
	return &productpb.Product{
		Id:          product.ID.String(),
		Name:        product.Name,
		Slug:        product.Slug,
		Description: product.Description,
		Price:       product.Price,
		WeightG:     product.WeightG,
		Thumbnail:   product.Thumbnail,
		// CreatedAt:   util.TimeToTimestamp(product.CreatedAt),
		// UpdatedAt:   util.TimeToTimestamp(product.UpdatedAt),
	}
}

func MapToProductDetail(product *model.Product) *productpb.ProductDetail {
	mappedProduct := MapToProduct(product)
	mappedCategories := MapProductToCategories(product)
	mappedImages := MapProductToImages(product)
	mappedReviews := MapProductToReviews(product)

	return &productpb.ProductDetail{
		Product:    mappedProduct,
		Images:     mappedImages,
		Categories: mappedCategories,
		Reviews:    mappedReviews,
	}
}

func MapToProductList(productList []model.Product, page int32, totalCount int32) *productpb.ProductList {
	var mappedProduct []*productpb.Product
	for _, product := range productList {
		mappedProduct = append(mappedProduct, MapToProduct(&product))
	}

	return &productpb.ProductList{
		Products:    mappedProduct,
		CurrentPage: page,
		TotalCount:  totalCount,
	}
}
