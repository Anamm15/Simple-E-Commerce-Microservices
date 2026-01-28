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

func (s *ShippingController) CalculateCost(ctx context.Context, request *shippingpb.ShippingCostRequest) (*shippingpb.ShippingCostResponse, error) {
	input := dto.CalculateShippingCostRequestDTO{
		OriginCityID:      request.OriginCityId,
		DestinationCityID: request.DestinationCityId,
		TotalWeightG:      request.TotalWeightG,
		CourierCode:       request.CourierCode,
		ServiceCode:       request.ServiceCode,
	}

	cost, err := s.ShippingService.CalculateCost(ctx, input)
	if err != nil {
		return nil, err
	}

	return cost, nil
}

func (s *ShippingController) InputTrackingNumber(ctx context.Context, request *shippingpb.InputTrackingRequest) (*shippingpb.ShipmentDetail, error) {
	input := dto.InputTrackingNumberRequestDTO{
		OrderID:        request.OrderId,
		TrackingNumber: request.TrackingNumber,
		CourierCode:    request.CourierCode,
	}

	shipment, err := s.ShippingService.InputTrackingNumber(ctx, input)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func (s *ShippingController) GetShipmentStatus(ctx context.Context, request *shippingpb.GetShipmentStatusRequest) (*shippingpb.ShipmentDetail, error) {
	shipment, err := s.ShippingService.GetShipmentStatus(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	return shipment, nil
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
