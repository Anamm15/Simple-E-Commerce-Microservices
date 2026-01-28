package dto

import (
	"inventory-service/internal/model"
	inventorypb "inventory-service/internal/pb/inventory"
	"inventory-service/internal/util"
)

type CreateStockRequestDTO struct {
	ProductID string
	Quantity  int32
}

type UpdateStockRequestDTO struct {
	ProductID string
	Quantity  int32
}

type ReserveAndReleaseStockRequestDTO struct {
	OrderID string
	Items   []*inventorypb.StockItem
}

func (dto *CreateStockRequestDTO) ToModel() *model.Inventory {
	productID, _ := util.StringToUUID(dto.ProductID)
	return &model.Inventory{
		ProductID: productID,
		Stock:     dto.Quantity,
	}
}
