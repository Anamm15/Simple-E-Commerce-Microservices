package dto

import (
	"product-service/internal/model"
	"product-service/pkg/util"
)

type SearchProductRequestDTO struct {
	SearchQuery string
	Category    string
	Sort        string
	MinPrice    int64
	MaxPrice    int64
	Page        int32
	Limit       int32
}

type CreateProductRequestDTO struct {
	Categories     []string
	Name           string
	Description    string
	Price          int64
	WeightG        int32
	Thumbnail      ProductImageRequestDTO
	AdditionalImgs []ProductImageRequestDTO
	InitialStock   int32
}

type UpdateProductRequestDTO struct {
	ProductID   string
	Categories  []*string
	Name        *string
	Description *string
	Price       *int64
	WeightG     *int32
}

type GetProductBatchRequestDTO struct {
	ProductIDs []string
}

type UpdateThumbnailProductRequest struct {
	ProductID string
	Image     ProductImageRequestDTO
}

func (r *CreateProductRequestDTO) ToModel(thumbnailUrl string, thumbnailPublicID string, additionalImages []model.Image) *model.Product {
	var categories []model.Category
	for _, category := range r.Categories {
		ID, _ := util.StringToUUID(category)
		categories = append(categories, model.Category{ID: ID})
	}

	return &model.Product{
		Name:              r.Name,
		Description:       r.Description,
		Price:             r.Price,
		WeightG:           r.WeightG,
		Thumbnail:         thumbnailUrl,
		ThumbnailPublicID: thumbnailPublicID,
		Categories:        categories,
		Images:            additionalImages,
	}
}
