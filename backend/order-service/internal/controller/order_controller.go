package controller

import (
	"context"

	"order-service/internal/dto"
	"order-service/internal/helper/enum"
	"order-service/internal/helper/mapper"
	orderpb "order-service/internal/pb/order"
	"order-service/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		Status:    enum.OrderStatus(request.Status),
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
	statusFilter, err := mapper.MapProtoOrderStatusToDomain(request.StatusFilter)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	input := dto.GetOrderHistoryRequestDTO{
		UserID:       request.UserId,
		Page:         request.Page,
		Limit:        request.Limit,
		StatusFilter: statusFilter,
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

func (c *orderController) Checkout(ctx context.Context, request *orderpb.CheckoutRequest) (*orderpb.CheckoutResponse, error) {
	input := dto.CheckoutRequestDTO{
		UserID:      request.UserId,
		Email:       request.Email,
		Address:     mapper.MapAddressProtoToDTO(request.Address),
		CourierCode: request.CourierCode,
		ServiceCode: request.ServiceCode,
		Products:    mapper.MapProductCheckoutProtoToDTO(request.Products),
	}

	return c.orderService.Create(ctx, input)
}

func (c *orderController) UpdateOrderStatus(ctx context.Context, request *orderpb.UpdateStatusRequest) (*orderpb.OrderUpdateResponse, error) {
	newStatus, err := mapper.MapProtoOrderStatusToDomain(request.NewStatus)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	input := dto.UpdateOrderStatusRequestDTO{
		OrderID:   request.OrderId,
		NewStatus: newStatus,
		Notes:     &request.Notes,
	}
	return c.orderService.Update(ctx, input)
}
