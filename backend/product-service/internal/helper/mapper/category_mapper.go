package mapper

import (
	"product-service/internal/model"
	productpb "product-service/internal/pb/product"
)

func MapProductToCategories(product *model.Product) []*productpb.Category {
	var categories []*productpb.Category
	for _, category := range product.Categories {
		categories = append(categories, &productpb.Category{
			Id:   category.ID.String(),
			Name: category.Name,
		})
	}
	return categories
}

func MapToCategory(category *model.Category) *productpb.Category {
	return &productpb.Category{
		Id:   category.ID.String(),
		Name: category.Name,
	}
}
