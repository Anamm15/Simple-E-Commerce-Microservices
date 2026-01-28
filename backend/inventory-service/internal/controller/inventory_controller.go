package controller

import (
	"context"

	"inventory-service/internal/dto"
	inventorypb "inventory-service/internal/pb/inventory"
	"inventory-service/internal/service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type InventoryController struct {
	InventoryService service.InventoryService
	inventorypb.UnimplementedInventoryServiceServer
}

func NewInventoryController(inventoryService service.InventoryService) *InventoryController {
	return &InventoryController{InventoryService: inventoryService}
}

func (c *InventoryController) CreateStock(ctx context.Context, request *inventorypb.CreateStockRequest) (*inventorypb.StockCount, error) {
	input := dto.CreateStockRequestDTO{
		ProductID: request.ProductId,
		Quantity:  request.Quantity,
	}

	createdInventory, err := c.InventoryService.CreateStock(ctx, input)
	if err != nil {
		return nil, err
	}

	return createdInventory, nil
}

func (c *InventoryController) CheckStock(ctx context.Context, request *inventorypb.CheckStockRequest) (*inventorypb.StockCount, error) {
	return c.InventoryService.CheckStock(ctx, request.ProductId)
}

func (c *InventoryController) ReserveStock(ctx context.Context, request *inventorypb.ReserveStockRequest) (*inventorypb.ReserveStockResponse, error) {
	input := dto.ReserveAndReleaseStockRequestDTO{
		OrderID: request.OrderId,
		Items:   request.Items,
	}

	reservedStock, err := c.InventoryService.ReserveStock(ctx, input)
	if err != nil {
		return nil, err
	}

	return reservedStock, nil
}

func (c *InventoryController) UpdateStock(ctx context.Context, request *inventorypb.UpdateStockRequest) (*inventorypb.StockCount, error) {
	input := dto.UpdateStockRequestDTO{
		ProductID: request.ProductId,
		Quantity:  request.Quantity,
	}
	updatedStock, err := c.InventoryService.UpdateStock(ctx, input)
	if err != nil {
		return nil, err
	}

	return updatedStock, nil
}

func (c *InventoryController) ReleaseStock(ctx context.Context, request *inventorypb.ReleaseStockRequest) (*emptypb.Empty, error) {
	input := dto.ReserveAndReleaseStockRequestDTO{
		OrderID: request.OrderId,
		Items:   request.Items,
	}

	err := c.InventoryService.ReleaseStock(ctx, input)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
