package dto

type UpdateStockRequestDTO struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required,min=1"`
}
