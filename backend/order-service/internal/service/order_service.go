package service

import (
	"context"
	"encoding/json"

	"order-service/internal/dto"
	"order-service/internal/helper"
	"order-service/internal/helper/enum"
	"order-service/internal/helper/mapper"
	"order-service/internal/model"
	inventorypb "order-service/internal/pb/inventory"
	orderpb "order-service/internal/pb/order"
	paymentpb "order-service/internal/pb/payment"
	productpb "order-service/internal/pb/product"
	shippingpb "order-service/internal/pb/shipping"
	"order-service/internal/repository"
	"order-service/internal/util"

	"github.com/google/uuid"
)

type OrderService interface {
	GetAll(ctx context.Context, request dto.AdminOrderFilterRequestDTO) (*orderpb.OrderList, error)
	GetByUserID(ctx context.Context, request dto.GetOrderHistoryRequestDTO) (*orderpb.OrderList, error)
	GetDetailOrder(ctx context.Context, orderID string) (*orderpb.OrderDetail, error)
	Create(ctx context.Context, request dto.CheckoutRequestDTO) (*orderpb.CheckoutResponse, error)
	Update(ctx context.Context, request dto.UpdateOrderStatusRequestDTO) (*orderpb.OrderUpdateResponse, error)
}

type orderService struct {
	orderRepo       repository.OrderRepository
	inventoryClient inventorypb.InventoryServiceClient
	paymentClient   paymentpb.PaymentServiceClient
	productClient   productpb.ProductServiceClient
	shippingClient  shippingpb.ShippingServiceClient
}

func NewOrderService(
	orderRepo repository.OrderRepository,
	inventoryClient inventorypb.InventoryServiceClient,
	paymentClient paymentpb.PaymentServiceClient,
	productClient productpb.ProductServiceClient,
	shippingClient shippingpb.ShippingServiceClient,
) OrderService {
	return &orderService{
		orderRepo:       orderRepo,
		inventoryClient: inventoryClient,
		paymentClient:   paymentClient,
		productClient:   productClient,
		shippingClient:  shippingClient,
	}
}

func (s *orderService) GetAll(ctx context.Context, request dto.AdminOrderFilterRequestDTO) (*orderpb.OrderList, error) {
	page, limit, _, filter := helper.QueryValidation(&request.Page, &request.Limit, nil, &request.Status)

	offset := helper.CalculateOffset(page, limit)
	orders, totalCount, err := s.orderRepo.GetAll(ctx, limit, offset, filter)
	if err != nil {
		return nil, err
	}

	var ordersIDs []string
	for _, order := range orders {
		ordersIDs = append(ordersIDs, order.ID.String())
	}

	return &orderpb.OrderList{
		Orders:      mapper.MapToOrderListResponse(orders),
		TotalCount:  totalCount,
		CurrentPage: request.Page,
	}, nil
}

func (s *orderService) GetByUserID(ctx context.Context, request dto.GetOrderHistoryRequestDTO) (*orderpb.OrderList, error) {
	page, limit, sort, filter := helper.QueryValidation(&request.Page, &request.Limit, &request.Sort, &request.StatusFilter)

	offset := helper.CalculateOffset(page, limit)
	userID, err := util.StringToUUID(request.UserID)
	if err != nil {
		return nil, err
	}

	orders, totalCount, err := s.orderRepo.GetByUserID(ctx, userID, limit, offset, sort, filter)
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

	shipment, err := s.shippingClient.GetShipment(ctx, &shippingpb.GetShipmentRequest{OrderId: orderID})
	if err != nil {
		return nil, err
	}

	return mapper.MapToOrderResponse(*order, shipment), nil
}

func (s *orderService) Create(ctx context.Context, request dto.CheckoutRequestDTO) (*orderpb.CheckoutResponse, error) {
	// 1. Get Detail product first
	var productIDs []string
	productQtyMap := make(map[string]int32)
	for _, product := range request.Products {
		productIDs = append(productIDs, product.ProductID)
		productQtyMap[product.ProductID] = product.Quantity
	}

	products, err := s.productClient.GetProductBatch(ctx, &productpb.GetProductBatchRequest{
		ProductIds: productIDs,
	})
	if err != nil {
		return nil, err
	}

	// 2. Extract field product for next process
	var totalAmount int64
	var totalWeightG int32
	stockItems := make([]*inventorypb.StockItem, 0)
	orderItems := make([]model.OrderItem, 0)
	for _, product := range products.Products {
		productIDParsed, _ := util.StringToUUID(product.Id)
		totalAmount += product.Price * int64(productQtyMap[product.Id])
		totalWeightG += product.WeightG * productQtyMap[product.Id]

		orderItems = append(orderItems, model.OrderItem{
			ProductID:                productIDParsed,
			ProductNameSnapshot:      product.Name,
			ProductPriceSnapshot:     product.Price,
			ProductThumbnailSnapshot: product.Thumbnail,
			ProductWeightSnapshot:    product.WeightG,
			Quantity:                 productQtyMap[product.Id],
		})

		stockItems = append(stockItems, &inventorypb.StockItem{
			ProductId: product.Id,
			Quantity:  productQtyMap[product.Id],
		})
	}

	// 3. Calculate final shipping fee
	shippingFee, err := s.shippingClient.CalculateFinalCost(ctx, &shippingpb.ShippingCostRequest{
		OriginCity:      "Jakarta",
		DestinationCity: request.Address.City,
		TotalWeightG:    totalWeightG,
		CourierCode:     request.CourierCode,
		ServiceCode:     request.ServiceCode,
	})
	if err != nil {
		return nil, err
	}

	// 4. Make an order data
	userIDParsed, err := util.StringToUUID(request.UserID)
	if err != nil {
		return nil, err
	}

	addressJSON, err := json.Marshal(request.Address)
	if err != nil {
		return nil, err
	}

	orderID := uuid.New()
	order := model.Order{
		ID:                      orderID,
		UserID:                  userIDParsed,
		TotalAmount:             totalAmount,
		ShippingCost:            shippingFee.Cost,
		ShippingAddressSnapshot: addressJSON,
		Items:                   orderItems,
		Status:                  enum.OrderStatusPending,
	}

	// 5. Reserve stock
	_, err = s.inventoryClient.ReserveStock(ctx, &inventorypb.ReserveStockRequest{
		OrderId: order.ID.String(),
		Items:   stockItems,
	})
	if err != nil {
		return nil, err
	}

	// 6. Request to payment gateway
	paymentResponse, err := s.paymentClient.CreatePayment(ctx, &paymentpb.CreatePaymentRequest{
		OrderId:             order.ID.String(),
		Amount:              order.TotalAmount + order.ShippingCost,
		CustomerName:        request.Address.Name,
		CustomerPhoneNumber: request.Address.Phone,
		CustomerEmail:       request.Email,
	})

	// 7. Save order
	err = s.orderRepo.Create(ctx, &order)
	if err != nil {
		return nil, err
	}

	return &orderpb.CheckoutResponse{
		OrderId:      order.ID.String(),
		PaymentUrl:   paymentResponse.RedirectUrl,
		PaymentToken: paymentResponse.Token,
	}, nil
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

	return &orderpb.OrderUpdateResponse{
		CurrentStatus: mapper.MapToProtoOrderStatus(order.Status),
	}, nil
}
