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

type ReserveStockRequestDTO struct {
	OrderID string
	Items   []*inventorypb.StockItem
}

type ReleaseStockRequestDTO struct {
	OrderID string
	Action  string
}

func (dto *CreateStockRequestDTO) ToModel() *model.Inventory {
	productID, _ := util.StringToUUID(dto.ProductID)
	return &model.Inventory{
		ProductID:  productID,
		TotalStock: dto.Quantity,
	}
}
