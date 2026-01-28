package service

import (
	"context"

	"shipping-service/internal/dto"
	"shipping-service/internal/helper/mapper"
	shippingpb "shipping-service/internal/pb/shipping"
	"shipping-service/internal/repository"
)

type ShippingService interface {
	Create(ctx context.Context, request dto.CreateShipmentRequestDTO) (*shippingpb.ShipmentDetail, error)
	CalculateCost(ctx context.Context, request dto.CalculateShippingCostRequestDTO) (*shippingpb.ShippingCostResponse, error)
	InputTrackingNumber(ctx context.Context, request dto.InputTrackingNumberRequestDTO) (*shippingpb.ShipmentDetail, error)
	GetShipmentStatus(ctx context.Context, orderID string) (*shippingpb.ShipmentDetail, error)
}

type shippingService struct {
	repository repository.ShippingRepository
}

func NewShippingService(repository repository.ShippingRepository) ShippingService {
	return &shippingService{repository: repository}
}

func (s *shippingService) Create(ctx context.Context, request dto.CreateShipmentRequestDTO) (*shippingpb.ShipmentDetail, error) {
	shipment := request.ToModel()
	err := s.repository.Create(ctx, shipment)
	if err != nil {
		return nil, err
	}

	return mapper.MapToDetailShipmentResponse(shipment), nil
}

func (s *shippingService) CalculateCost(ctx context.Context, request dto.CalculateShippingCostRequestDTO) (*shippingpb.ShippingCostResponse, error) {
	cost := &shippingpb.ShippingCostResponse{
		CourierCode:   "JNT",
		ServiceCode:   "78972",
		Cost:          50000,
		EstimatedTime: "5",
	}

	return cost, nil
}

func (s *shippingService) InputTrackingNumber(ctx context.Context, request dto.InputTrackingNumberRequestDTO) (*shippingpb.ShipmentDetail, error) {
	return nil, nil
}

func (s *shippingService) GetShipmentStatus(ctx context.Context, orderID string) (*shippingpb.ShipmentDetail, error) {
	return nil, nil
}
