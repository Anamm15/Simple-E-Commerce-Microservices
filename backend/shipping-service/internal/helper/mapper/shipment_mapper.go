package mapper

import (
	"shipping-service/internal/model"
	shippingpb "shipping-service/internal/pb/shipping"
)

func MapToDetailShipmentResponse(shipment *model.Shipment) *shippingpb.ShipmentDetail {
	orderID := shipment.OrderID.String()
	return &shippingpb.ShipmentDetail{
		OrderId:        orderID,
		TrackingNumber: *shipment.TrackingNumber,
		CourierCode:    shipment.CourierCode,
		ServiceCode:    shipment.ServiceCode,
		ShippingCost:   shipment.ShippingCost,
	}
}

func MapToListDetailShipmentResponse(shipments []model.Shipment) []*shippingpb.ShipmentDetail {
	var response []*shippingpb.ShipmentDetail
	for _, shipment := range shipments {
		response = append(response, MapToDetailShipmentResponse(&shipment))
	}
	return response
}
