package controller

import (
	"context"

	"shipping-service/internal/dto"
	shippingpb "shipping-service/internal/pb/shipping"
	"shipping-service/internal/service"
)

type ShippingController struct {
	shippingpb.UnimplementedShippingServiceServer
	ShippingService service.ShippingService
}

func NewShippingController(shippingService service.ShippingService) *ShippingController {
	return &ShippingController{ShippingService: shippingService}
}

func (s *ShippingController) OfferShippingCost(ctx context.Context, request *shippingpb.ShippingCostOfferRequest) (*shippingpb.ShippingCostOfferResponse, error) {
	input := dto.OfferShippingCostRequestDTO{
		OriginCity:      request.OriginCity,
		DestinationCity: request.DestinationCity,
		TotalWeightG:    request.TotalWeightG,
	}

	return s.ShippingService.OfferShippingCost(ctx, input)
}

func (s *ShippingController) CalculateFinalCost(ctx context.Context, request *shippingpb.ShippingCostRequest) (*shippingpb.ShippingCostResponse, error) {
	input := dto.CalculateShippingCostRequestDTO{
		OriginCity:      request.OriginCity,
		DestinationCity: request.DestinationCity,
		TotalWeightG:    request.TotalWeightG,
		CourierCode:     request.CourierCode,
		ServiceCode:     request.ServiceCode,
	}

	cost, err := s.ShippingService.CalculateCost(ctx, input)
	if err != nil {
		return nil, err
	}

	return cost, nil
}

func (s *ShippingController) GetShipment(ctx context.Context, request *shippingpb.GetShipmentRequest) (*shippingpb.ShipmentDetail, error) {
	shipment, err := s.ShippingService.GetShipment(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func (s *ShippingController) GetBatchShipment(ctx context.Context, request *shippingpb.GetBatchShipmentsRequest) (*shippingpb.GetBatchShipmentsResponse, error) {
	shipments, err := s.ShippingService.GetBatchShipments(ctx, request.ShippingIds)
	if err != nil {
		return nil, err
	}

	return shipments, nil
}

func (s *ShippingController) CreateShipment(ctx context.Context, request *shippingpb.CreateShipmentRequest) (*shippingpb.ShipmentDetail, error) {
	input := dto.CreateShipmentRequestDTO{
		OrderID:        request.OrderId,
		CourierCode:    request.CourierCode,
		ServiceCode:    request.ServiceCode,
		TrackingNumber: request.TrackingNumber,
		ShippingCost:   request.ShippingCost,
	}

	shipment, err := s.ShippingService.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}
