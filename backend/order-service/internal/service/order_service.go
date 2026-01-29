package service

import (
	"context"

	"order-service/internal/dto"
	"order-service/internal/helper"
	"order-service/internal/helper/mapper"
	inventorypb "order-service/internal/pb/inventory"
	notificationpb "order-service/internal/pb/notification"
	orderpb "order-service/internal/pb/order"
	paymentpb "order-service/internal/pb/payment"
	productpb "order-service/internal/pb/product"
	shippingpb "order-service/internal/pb/shipping"
	"order-service/internal/repository"
	"order-service/internal/util"
)

type OrderService interface {
	GetAll(ctx context.Context, request dto.AdminOrderFilterRequestDTO) (*orderpb.OrderList, error)
	GetByUserID(ctx context.Context, request dto.GetOrderHistoryRequestDTO) (*orderpb.OrderList, error)
	GetDetailOrder(ctx context.Context, orderID string) (*orderpb.OrderDetail, error)
	Create(ctx context.Context, request dto.CheckoutRequestDTO) (*orderpb.CheckoutResponse, error)
	Update(ctx context.Context, request dto.UpdateOrderStatusRequestDTO) (*orderpb.OrderUpdateResponse, error)
}

type orderService struct {
	orderRepo          repository.OrderRepository
	inventoryClient    inventorypb.InventoryServiceClient
	paymentClient      paymentpb.PaymentServiceClient
	productClient      productpb.ProductServiceClient
	shippingClient     shippingpb.ShippingServiceClient
	notificationClient notificationpb.NotificationServiceClient
}

func NewOrderService(
	orderRepo repository.OrderRepository,
	inventoryClient inventorypb.InventoryServiceClient,
	paymentClient paymentpb.PaymentServiceClient,
	productClient productpb.ProductServiceClient,
	shippingClient shippingpb.ShippingServiceClient,
	notificationClient notificationpb.NotificationServiceClient,
) OrderService {
	return &orderService{
		orderRepo:          orderRepo,
		inventoryClient:    inventoryClient,
		paymentClient:      paymentClient,
		productClient:      productClient,
		shippingClient:     shippingClient,
		notificationClient: notificationClient,
	}
}

func (s *orderService) GetAll(ctx context.Context, request dto.AdminOrderFilterRequestDTO) (*orderpb.OrderList, error) {
	offset := helper.CalculateOffset(&request.Page, &request.Limit)
	orders, totalCount, err := s.orderRepo.GetAll(ctx, &request.Limit, &offset, nil, &request.Status)
	if err != nil {
		return nil, err
	}

	return &orderpb.OrderList{
		Orders:      mapper.MapToOrderListResponse(orders),
		TotalCount:  totalCount,
		CurrentPage: request.Page,
	}, nil
}

func (s *orderService) GetByUserID(ctx context.Context, request dto.GetOrderHistoryRequestDTO) (*orderpb.OrderList, error) {
	offset := helper.CalculateOffset(&request.Page, &request.Limit)
	userID, err := util.StringToUUID(request.UserID)
	if err != nil {
		return nil, err
	}

	orders, totalCount, err := s.orderRepo.GetByUserID(ctx, userID, &request.Limit, &offset, &request.Sort, &request.StatusFilter)
	if err != nil {
		return nil, err
	}

	return &orderpb.OrderList{
		Orders:      mapper.MapToOrderListResponse(orders),
		TotalCount:  totalCount,
		CurrentPage: request.Page,
	}, nil
}

func (s *orderService) GetDetailOrder(ctx context.Context, orderID string) (*orderpb.OrderDetail, error) {
	orderIDParsed, err := util.StringToUUID(orderID)
	if err != nil {
		return nil, err
	}

	order, err := s.orderRepo.GetDetailOrder(ctx, orderIDParsed)
	if err != nil {
		return nil, err
	}

	return mapper.MapToOrderResponse(*order), nil
}

func (s *orderService) Create(ctx context.Context, request dto.CheckoutRequestDTO) (*orderpb.CheckoutResponse, error) {
	return nil, nil
}

func (s *orderService) Update(ctx context.Context, request dto.UpdateOrderStatusRequestDTO) (*orderpb.OrderUpdateResponse, error) {
	orderID, err := util.StringToUUID(request.OrderID)
	if err != nil {
		return nil, err
	}

	order, err := s.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	order.Status = request.NewStatus
	if request.Notes != nil {
		order.Notes = request.Notes
	}

	err = s.orderRepo.Update(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orderpb.OrderUpdateResponse{CurrentStatus: order.Status}, nil
}
