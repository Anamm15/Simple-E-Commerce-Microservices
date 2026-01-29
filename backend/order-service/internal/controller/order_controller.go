package controller

import (
	"context"

	"order-service/internal/dto"
	orderpb "order-service/internal/pb/order"
	"order-service/internal/service"
)

type orderController struct {
	orderpb.UnimplementedOrderServiceServer
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) orderpb.OrderServiceServer {
	return &orderController{orderService: orderService}
}

func (c *orderController) GetAllOrders(ctx context.Context, request *orderpb.AdminFilterRequest) (*orderpb.OrderList, error) {
	input := dto.AdminOrderFilterRequestDTO{
		Page:      request.Page,
		Limit:     request.Limit,
		UserID:    request.UserId,
		Status:    request.Status,
		DateStart: request.DateStart,
		DateEnd:   request.DateEnd,
	}

	orders, err := c.orderService.GetAll(ctx, input)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *orderController) GetOrderHistory(ctx context.Context, request *orderpb.GetOrderHistoryRequest) (*orderpb.OrderList, error) {
	input := dto.GetOrderHistoryRequestDTO{
		UserID:       request.UserId,
		Page:         request.Page,
		Limit:        request.Limit,
		StatusFilter: request.StatusFilter,
		Sort:         request.Sort,
	}

	orders, err := c.orderService.GetByUserID(ctx, input)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *orderController) GetOrderDetail(ctx context.Context, request *orderpb.GetOrderDetailRequest) (*orderpb.OrderDetail, error) {
	return c.orderService.GetDetailOrder(ctx, request.OrderId)
}

func (c *orderController) CreateOrder(ctx context.Context, request *orderpb.CheckoutRequest) (*orderpb.CheckoutResponse, error) {
	input := dto.CheckoutRequestDTO{
		UserID:      request.UserId,
		Address:     request.Address,
		CourierCode: request.CourierCode,
		ServiceCode: request.ServiceCode,
		ProductIDs:  request.ProductIds,
	}

	return c.orderService.Create(ctx, input)
}

func (c *orderController) UpdateOrderStatus(ctx context.Context, request *orderpb.UpdateStatusRequest) (*orderpb.OrderUpdateResponse, error) {
	input := dto.UpdateOrderStatusRequestDTO{
		OrderID:   request.OrderId,
		NewStatus: request.NewStatus,
		Notes:     &request.Notes,
	}

	return c.orderService.Update(ctx, input)
}
