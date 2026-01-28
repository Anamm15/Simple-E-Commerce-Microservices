package mapper

import (
	"inventory-service/internal/model"
	inventorypb "inventory-service/internal/pb/inventory"
)

func MapToInventoryStockCount(inventory *model.Inventory) *inventorypb.StockCount {
	return &inventorypb.StockCount{
		ProductId:  inventory.ProductID.String(),
		TotalStock: inventory.Stock,
	}
}
