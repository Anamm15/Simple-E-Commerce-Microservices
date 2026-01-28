package dto

type AddCartItemRequestDTO struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required,min=1"`
}

type RemoveCartItemRequestDTO struct {
	ProductID string `json:"product_id" binding:"required"`
}
