package dto

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
