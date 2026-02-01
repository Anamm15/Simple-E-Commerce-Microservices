package service

import (
	"context"

	"shipping-service/internal/dto"
	"shipping-service/internal/helper/algorithm"
	"shipping-service/internal/helper/mapper"
	shippingpb "shipping-service/internal/pb/shipping"
	"shipping-service/internal/repository"
	"shipping-service/internal/util"
)

type ShippingService interface {
	OfferShippingCost(ctx context.Context, request dto.OfferShippingCostRequestDTO) (*shippingpb.ShippingCostOfferResponse, error)
	Create(ctx context.Context, request dto.CreateShipmentRequestDTO) (*shippingpb.ShipmentDetail, error)
	CalculateCost(ctx context.Context, request dto.CalculateShippingCostRequestDTO) (*shippingpb.ShippingCostResponse, error)
	GetShipment(ctx context.Context, orderID string) (*shippingpb.ShipmentDetail, error)
}

type shippingService struct {
	shipmentRepo repository.ShippingRepository
}

func NewShippingService(shipmentRepo repository.ShippingRepository) ShippingService {
	return &shippingService{shipmentRepo: shipmentRepo}
}

func (s *shippingService) OfferShippingCost(ctx context.Context, request dto.OfferShippingCostRequestDTO) (*shippingpb.ShippingCostOfferResponse, error) {
	options := algorithm.CalculateShipping(request.OriginCity, request.DestinationCity, request.TotalWeightG)
	if len(options) == 0 {
		return &shippingpb.ShippingCostOfferResponse{
			Options: []*shippingpb.ShippingOption{},
		}, nil
	}

	pbOptions := make([]*shippingpb.ShippingOption, 0, len(options))

	for _, opt := range options {
		pbOptions = append(pbOptions, &shippingpb.ShippingOption{
			Courier:       opt.Courier,
			Service:       opt.Service,
			Cost:          opt.Cost,
			EstimatedDays: opt.EstimatedDays,
		})
	}

	return &shippingpb.ShippingCostOfferResponse{
		Options: pbOptions,
	}, nil
}

func (s *shippingService) Create(ctx context.Context, request dto.CreateShipmentRequestDTO) (*shippingpb.ShipmentDetail, error) {
	shipment := request.ToModel()
	err := s.shipmentRepo.Create(ctx, shipment)
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

func (s *shippingService) GetShipment(ctx context.Context, orderID string) (*shippingpb.ShipmentDetail, error) {
	orderIDParsed, err := util.StringToUUID(orderID)
	if err != nil {
		return &shippingpb.ShipmentDetail{}, nil
	}

	shipment, err := s.shipmentRepo.GetByOrderID(ctx, orderIDParsed)
	return mapper.MapToDetailShipmentResponse(shipment), nil
}
