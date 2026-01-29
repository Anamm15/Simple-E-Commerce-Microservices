package dto

type UpdateStockRequestDTO struct {
	ProductID  string `json:"product_id" binding:"required"`
	UpdateType string `json:"update_type" binding:"required,oneof=ADD SET"` // ADD untuk nambah, SET untuk opname
	Quantity   int32  `json:"quantity" binding:"required,min=1"`
}
