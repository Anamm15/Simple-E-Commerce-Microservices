package controller

import (
	"context"

	"user-service/internal/dto"
	"user-service/internal/service"

	userpb "user-service/internal/pb/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userController struct {
	userpb.UnimplementedUserServiceServer
	userService    service.UserService
	addressService service.AddressService
}

func NewUserController(userService service.UserService, addressService service.AddressService) *userController {
	return &userController{
		userService:    userService,
		addressService: addressService,
	}
}

func (c *userController) GetUserProfile(ctx context.Context, req *userpb.GetUserProfileRequest) (*userpb.UserProfile, error) {
	user, err := c.userService.GetByID(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	return user, nil
}

func (c *userController) UpdateUserProfile(ctx context.Context, req *userpb.UpdateUserProfileRequest) (*userpb.UserProfile, error) {
	input := dto.UpdateUserProfileRequestDTO{
		UserID:      req.UserId,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
	}

	updatedUser, err := c.userService.Update(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update profile: %v", err)
	}

	return updatedUser, nil
}

func (c *userController) AddAddress(ctx context.Context, req *userpb.AddAddressRequest) (*userpb.AddressResponse, error) {
	input := dto.AddAddressRequestDTO{
		UserID:    req.UserId,
		Street:    req.Street,
		City:      req.City,
		Province:  req.Province,
		ZipCode:   req.ZipCode,
		IsPrimary: req.IsPrimary,
	}

	newAddress, err := c.addressService.Create(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add address: %v", err)
	}

	return newAddress, nil
}

func (c *userController) GetAddress(ctx context.Context, req *userpb.GetAddressRequest) (*userpb.AddressDetail, error) {
	address, err := c.addressService.GetByID(ctx, req.AddressId, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "address not found: %v", err)
	}

	return address, nil
}

func (c *userController) UpdateAddress(ctx context.Context, req *userpb.UpdateAddressRequest) (*userpb.AddressDetail, error) {
	input := dto.UpdateAddressRequestDTO{
		AddressID: req.AddressId,
		UserID:    req.UserId,
		Street:    req.Street,
		City:      req.City,
		Province:  req.Province,
		ZipCode:   req.ZipCode,
		IsPrimary: &req.IsPrimary,
	}

	updatedAddress, err := c.addressService.Update(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update address: %v", err)
	}

	return updatedAddress, nil
}

func (c *userController) DeleteAddress(ctx context.Context, req *userpb.DeleteAddressRequest) (*emptypb.Empty, error) {
	err := c.addressService.Delete(ctx, req.AddressId, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete address: %v", err)
	}

	return &emptypb.Empty{}, nil
}
