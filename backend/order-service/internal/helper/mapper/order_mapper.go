package mapper

import (
	"encoding/json"

	"order-service/internal/model"
	orderpb "order-service/internal/pb/order"
	"order-service/internal/util"
)

func mapAddressSnapshot(raw string) (*orderpb.AddressSnapshot, error) {
	var addr orderpb.AddressSnapshot
	if err := json.Unmarshal([]byte(raw), &addr); err != nil {
		return nil, err
	}
	return &addr, nil
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

func MapToOrderResponse(order model.Order) *orderpb.OrderDetail {
	address, _ := mapAddressSnapshot(order.ShippingAddressSnapshot)

	items := make([]*orderpb.OrderItem, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, MapToOrderItem(&item))
	}

	return &orderpb.OrderDetail{
		Id:           util.UUIDToString(order.ID),
		UserId:       util.UUIDToString(order.UserID),
		Status:       order.Status,
		TotalAmount:  order.TotalAmount,
		ShippingCost: order.ShippingCost,

		ShippingAddress: address,

		Items: items,
		// CreatedAt: timestamppb.New(order.CreatedAt),
	}
}

func MapToOrderListResponse(orders []model.Order) []*orderpb.OrderDetail {
	res := make([]*orderpb.OrderDetail, 0, len(orders))
	for _, order := range orders {
		res = append(res, MapToOrderResponse(order))
	}
	return res
}
