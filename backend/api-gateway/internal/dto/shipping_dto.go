package dto

type OfferShippingCostRequestDTO struct {
	OriginCity      string `json:"origin_city" binding:"required"`
	DestinationCity string `json:"destination_city" binding:"required"`
	TotalWeightG    int32  `json:"total_weight_g" binding:"required,min=1"`
}
