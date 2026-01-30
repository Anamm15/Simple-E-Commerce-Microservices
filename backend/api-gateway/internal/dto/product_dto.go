package dto

type SearchProductRequestDTO struct {
	SearchQuery  string `form:"search"`
	CategorySlug string `form:"category"`
	MinPrice     int64  `form:"min_price" binding:"omitempty,min=0"`
	MaxPrice     int64  `form:"max_price" binding:"omitempty,min=0,gtfield=MinPrice"`
	Page         int32  `form:"page,default=1"`
	Limit        int32  `form:"limit,default=10"`
}

type CreateProductRequestDTO struct {
	CategoryID          string   `json:"category_id" binding:"required"`
	Name                string   `json:"name" binding:"required,min=3"`
	Description         string   `json:"description" binding:"required"`
	Price               int64    `json:"price" binding:"required,min=100"`
	WeightG             int32    `json:"weight_g" binding:"required,min=1"`
	Thumbnail           string   `json:"image_url" binding:"required,url"`
	AdditionalImageURLs []string `json:"additional_image_urls"`
	InitialStock        int32    `json:"initial_stock" binding:"required,min=0"`
}

type UpdateProductRequestDTO struct {
	CategoryID          string   `json:"category_id" binding:"required"`
	Name                string   `json:"name" binding:"required"`
	Description         string   `json:"description" binding:"required"`
	Price               int64    `json:"price" binding:"required,min=100"`
	WeightG             int32    `json:"weight_g" binding:"required,min=1"`
	ImageURL            string   `json:"image_url" binding:"required,url"`
	AdditionalImageURLs []string `json:"additional_image_urls"`
}

type GetProductBatchRequestDTO struct {
	ProductIDs []string `json:"product_ids" binding:"required,dive,uuid"`
}
