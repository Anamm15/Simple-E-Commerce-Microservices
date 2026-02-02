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
	Categories   []string `form:"category" binding:"required"`
	Name         string   `form:"name" binding:"required,min=3"`
	Description  string   `form:"description" binding:"required"`
	Price        int64    `form:"price" binding:"required,min=100"`
	WeightG      int32    `form:"weight_g" binding:"required,min=1"`
	InitialStock int32    `form:"initial_stock" binding:"required,min=0"`
}

type UpdateProductRequestDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	WeightG     int32  `json:"weight_g"`
}

type GetProductBatchRequestDTO struct {
	ProductIDs []string `json:"product_ids" binding:"required,dive,uuid"`
}
