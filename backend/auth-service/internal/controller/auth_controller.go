package controller

import (
	"context"

	"auth-service/internal/dto"
	authpb "auth-service/internal/pb/auth"
	"auth-service/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type authController struct {
	authpb.UnimplementedAuthServiceServer
	service service.AuthService
}

func NewAuthController(service service.AuthService) *authController {
	return &authController{service: service}
}

func (c *authController) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	input := dto.RegisterRequestDTO{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	data, err := c.service.Register(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register user: %v", err)
	}

	return &authpb.RegisterResponse{
		Id: data.ID.String(),
	}, nil
}

func (c *authController) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	input := dto.LoginRequestDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := c.service.Login(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials: %v", err)
	}

	return &authpb.LoginResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiresIn:    result.ExpiresIn,
	}, nil
}

func (c *authController) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.TokenResponse, error) {
	input := dto.RefreshTokenRequestDTO{
		UserID: req.UserId,
		Email:  req.Email,
		Role:   req.Role,
	}

	result, err := c.service.RefreshToken(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to refresh token: %v", err)
	}

	return &authpb.TokenResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiresIn:    result.ExpiresIn,
	}, nil
}

func (c *authController) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*emptypb.Empty, error) {
	input := dto.ChangePasswordRequestDTO{
		UserID:      req.UserId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	err := c.service.ChangePassword(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to change password: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (c *authController) ResetPassword(ctx context.Context, req *authpb.ResetPasswordRequest) (*emptypb.Empty, error) {
	input := dto.ResetPasswordRequestDTO{
		ResetToken:  req.ResetToken,
		NewPassword: req.NewPassword,
	}

	err := c.service.ResetPassword(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to reset password: %v", err)
	}

	return &emptypb.Empty{}, nil
}
