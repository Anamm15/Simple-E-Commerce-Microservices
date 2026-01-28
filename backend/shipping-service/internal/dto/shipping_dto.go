package dto

import (
	"shipping-service/internal/model"
	"shipping-service/internal/util"
)

type CalculateShippingCostRequestDTO struct {
	OriginCityID      string `json:"origin_city_id" binding:"required"`
	DestinationCityID string `json:"destination_city_id" binding:"required"`
	TotalWeightG      int32  `json:"total_weight_g" binding:"required,min=1"`
	CourierCode       string `json:"courier_code" binding:"required"`
	ServiceCode       string `json:"service_code"`
}

type InputTrackingNumberRequestDTO struct {
	OrderID        string `json:"order_id" binding:"required"`
	TrackingNumber string `json:"tracking_number" binding:"required"`
	CourierCode    string `json:"courier_code" binding:"required"`
}

type CreateShipmentRequestDTO struct {
	OrderID        string `json:"order_id" binding:"required"`
	CourierCode    string `json:"courier_code" binding:"required"`
	ServiceCode    string `json:"service_code" binding:"required"`
	TrackingNumber string `json:"tracking_number" binding:"required"`
	ShippingCost   int64  `json:"shipping_cost" binding:"required"`
}

func (dto *CreateShipmentRequestDTO) ToModel() *model.Shipment {
	orderID, _ := util.StringToUUID(dto.OrderID)
	return &model.Shipment{
		OrderID:        orderID,
		CourierCode:    dto.CourierCode,
		ServiceCode:    dto.ServiceCode,
		TrackingNumber: dto.TrackingNumber,
		ShippingCost:   dto.ShippingCost,
	}
}
