package mapper

import (
	"encoding/json"
	"fmt"

	"order-service/internal/dto"
	"order-service/internal/helper/enum"
	"order-service/internal/model"
	orderpb "order-service/internal/pb/order"
	shippingpb "order-service/internal/pb/shipping"
	"order-service/internal/util"

	"gorm.io/datatypes"
)

func MapProtoOrderStatusToDomain(status orderpb.OrderStatus) (enum.OrderStatus, error) {
	switch status {
	case orderpb.OrderStatus_PENDING:
		return enum.OrderStatusPending, nil
	case orderpb.OrderStatus_PAID:
		return enum.OrderStatusPaid, nil
	case orderpb.OrderStatus_PROCESSED:
		return enum.OrderStatusProcessed, nil
	case orderpb.OrderStatus_SHIPPED:
		return enum.OrderStatusShipped, nil
	case orderpb.OrderStatus_DELIVERED:
		return enum.OrderStatusDelivered, nil
	case orderpb.OrderStatus_CANCELLED:
		return enum.OrderStatusCancelled, nil
	case orderpb.OrderStatus_FAILED:
		return enum.OrderStatusFailed, nil
	default:
		return "", fmt.Errorf("unknown order status: %v", status)
	}
}

func MapToProtoOrderStatus(status enum.OrderStatus) orderpb.OrderStatus {
	switch status {
	case enum.OrderStatusPending:
		return orderpb.OrderStatus_PENDING
	case enum.OrderStatusPaid:
		return orderpb.OrderStatus_PAID
	case enum.OrderStatusProcessed:
		return orderpb.OrderStatus_PROCESSED
	case enum.OrderStatusShipped:
		return orderpb.OrderStatus_SHIPPED
	case enum.OrderStatusDelivered:
		return orderpb.OrderStatus_DELIVERED
	case enum.OrderStatusCancelled:
		return orderpb.OrderStatus_CANCELLED
	case enum.OrderStatusFailed:
		return orderpb.OrderStatus_FAILED
	default:
		return orderpb.OrderStatus_FAILED
	}
}

func mapAddressSnapshot(addressJson datatypes.JSON) (*orderpb.AddressSnapshot, error) {
	address := &orderpb.AddressSnapshot{}
	err := json.Unmarshal(addressJson, address)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func MapAddressProtoToDTO(address *orderpb.AddressSnapshot) dto.AddressRequestDTO {
	return dto.AddressRequestDTO{
		Street:   address.Street,
		City:     address.City,
		ZipCode:  address.ZipCode,
		Province: address.Province,
		Name:     address.ReceiverName,
		Phone:    address.ReceiverPhone,
	}
}

func MapToOrderItem(item *model.OrderItem) *orderpb.OrderItem {
	return &orderpb.OrderItem{
		Id:                       util.UUIDToString(item.ID),
		ProductId:                util.UUIDToString(item.ProductID),
		Quantity:                 item.Quantity,
		ProductPriceSnapshot:     item.ProductPriceSnapshot,
		ProductNameSnapshot:      item.ProductNameSnapshot,
		ProductThumbnailSnapshot: item.ProductThumbnailSnapshot,
		ProductWeightSnapshot:    item.ProductWeightSnapshot,
	}
}

func MapToOrderResponse(
	order model.Order,
	shipment *shippingpb.ShipmentDetail,
) *orderpb.OrderDetail {
	address, _ := mapAddressSnapshot(order.ShippingAddressSnapshot)

	items := make([]*orderpb.OrderItem, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, MapToOrderItem(&item))
	}

	resp := &orderpb.OrderDetail{
		Id:              util.UUIDToString(order.ID),
		UserId:          util.UUIDToString(order.UserID),
		Status:          MapToProtoOrderStatus(order.Status),
		TotalAmount:     order.TotalAmount,
		ShippingCost:    order.ShippingCost,
		ShippingAddress: address,
		Items:           items,
	}

	if shipment != nil {
		resp.ShippingCourier = shipment.CourierCode
		resp.ShippingService = shipment.ServiceCode
		resp.TrackingNumber = shipment.TrackingNumber
		resp.ShippingStatus = shipment.Status
	}

	return resp
}

func MapToOrderListResponse(orders []model.Order) []*orderpb.OrderDetail {
	res := make([]*orderpb.OrderDetail, 0, len(orders))
	for _, order := range orders {
		res = append(res, MapToOrderResponse(order, nil))
	}
	return res
}
