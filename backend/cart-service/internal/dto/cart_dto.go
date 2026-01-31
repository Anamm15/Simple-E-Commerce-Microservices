package dto

import (
	"cart-service/internal/model"
	"cart-service/pkg/util"
)

type AddCartItemRequestDTO struct {
	UserID    string `json:"user_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required,min=1"`
}

type RemoveCartItemRequestDTO struct {
	UserID    string `json:"user_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
}

type UpdateCartItemRequestDTO struct {
	UserID    string `json:"user_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required,min=1"`
}

func (r *AddCartItemRequestDTO) ToModel() *model.CartItem {
	userID, _ := util.StringToUUID(r.UserID)
	productID, _ := util.StringToUUID(r.ProductID)

	return &model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  r.Quantity,
	}
}
