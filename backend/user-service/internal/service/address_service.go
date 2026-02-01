package service

import (
	"context"

	"user-service/internal/dto"
	"user-service/internal/helper/mapper"
	userpb "user-service/internal/pb/user"
	"user-service/internal/repository"

	"github.com/google/uuid"
)

type AddressService interface {
	GetAll(ctx context.Context, addressID string, userID string) ([]userpb.AddressDetail, error)
	GetByID(ctx context.Context, addressID string, userID string) (*userpb.AddressDetail, error)
	Create(ctx context.Context, req dto.AddAddressRequestDTO) (*userpb.AddressResponse, error)
	Update(ctx context.Context, req dto.UpdateAddressRequestDTO) (*userpb.AddressDetail, error)
	Delete(ctx context.Context, addressID string, userID string) error
}

type addressService struct {
	addressRepo repository.AddressRepository
}

func NewAddressService(addressRepo repository.AddressRepository) AddressService {
	return &addressService{addressRepo: addressRepo}
}

func (s *addressService) GetAll(ctx context.Context, addressID string, userID string) ([]userpb.AddressDetail, error) {
	userIDParsed, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	addresses, err := s.addressRepo.GetByUserID(ctx, userIDParsed)
	if err != nil {
		return nil, err
	}

	userAddresses := mapper.AddressDetailListResponseMapper(addresses)
	return userAddresses, nil
}

func (s *addressService) GetByID(ctx context.Context, addressID string, userID string) (*userpb.AddressDetail, error) {
	addressIDParsed, err := uuid.Parse(addressID)
	if err != nil {
		return &userpb.AddressDetail{}, err
	}

	address, err := s.addressRepo.GetByID(ctx, addressIDParsed)
	if err != nil {
		return &userpb.AddressDetail{}, err
	}

	addressDetail := mapper.AddressDetailResponseMapper(address)
	return addressDetail, nil
}

func (s *addressService) Create(ctx context.Context, req dto.AddAddressRequestDTO) (*userpb.AddressResponse, error) {
	address := req.ToModel()
	err := s.addressRepo.Create(ctx, address)
	if err != nil {
		return &userpb.AddressResponse{}, err
	}

	addressResponse := mapper.AddressResponseMapper(address)
	return addressResponse, nil
}

func (s *addressService) Update(ctx context.Context, req dto.UpdateAddressRequestDTO) (*userpb.AddressDetail, error) {
	userIDParsed, err := uuid.Parse(req.UserID)
	if err != nil {
		return &userpb.AddressDetail{}, err
	}

	addressIDParsed, err := uuid.Parse(req.AddressID)
	if err != nil {
		return &userpb.AddressDetail{}, err
	}

	address, err := s.addressRepo.GetByID(ctx, addressIDParsed)
	if err != nil {
		return &userpb.AddressDetail{}, err
	}

	if address.UserID != userIDParsed {
		return &userpb.AddressDetail{}, err
	}

	if req.Street != "" {
		address.Street = req.Street
	}

	if req.City != "" {
		address.City = req.City
	}

	if req.Province != "" {
		address.Province = req.Province
	}

	if req.ZipCode != "" {
		address.ZipCode = req.ZipCode
	}

	if req.IsPrimary != nil {
		address.IsPrimary = *req.IsPrimary
	}

	err = s.addressRepo.Update(ctx, address)
	if err != nil {
		return &userpb.AddressDetail{}, err
	}

	addressDetail := mapper.AddressDetailResponseMapper(address)
	return addressDetail, nil
}

func (s *addressService) Delete(ctx context.Context, addressID string, userID string) error {
	userIDParsed, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	addressIDParsed, err := uuid.Parse(addressID)
	if err != nil {
		return err
	}

	return s.addressRepo.Delete(ctx, addressIDParsed, userIDParsed)
}
