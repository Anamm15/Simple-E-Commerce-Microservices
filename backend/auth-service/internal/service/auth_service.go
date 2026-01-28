package service

import (
	"context"

	"auth-service/internal/dto"
	"auth-service/internal/helper"
	"auth-service/internal/helper/constant"
	"auth-service/internal/helper/enum"
	"auth-service/internal/repository"
	"auth-service/internal/util"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequestDTO) (dto.AccountResponseDTO, error)
	Login(ctx context.Context, req dto.LoginRequestDTO) (dto.LoginResponseDTO, error)
	RefreshToken(ctx context.Context, req dto.RefreshTokenRequestDTO) (dto.TokenResponseDTO, error)
	ChangePassword(ctx context.Context, req dto.ChangePasswordRequestDTO) error
	ResetPassword(ctx context.Context, req dto.ResetPasswordRequestDTO) error
}

type authService struct {
	authRepo repository.AccountRepository
}

func NewAuthService(authRepo repository.AccountRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequestDTO) (dto.AccountResponseDTO, error) {
	account := req.ToModel()
	err := s.authRepo.Create(ctx, account)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	return *helper.MapAccountToAccountResponseDTO(account), nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequestDTO) (dto.LoginResponseDTO, error) {
	account, err := s.authRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	err = util.ComparePassword(account.Password, req.Password)
	if err != nil {
		return dto.LoginResponseDTO{}, constant.ErrInvalidCredentials
	}

	accessToken, err := util.GenerateJWT(account.ID, account.Email, account.Role)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	refreshToken, err := util.GenerateRandomString(32)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	return dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    24 * 60 * 60,
	}, nil
}

func (s *authService) RefreshToken(ctx context.Context, req dto.RefreshTokenRequestDTO) (dto.TokenResponseDTO, error) {
	userID, err := util.StringToUUID(req.UserID)
	if err != nil {
		return dto.TokenResponseDTO{}, err
	}

	accessToken, err := util.GenerateJWT(userID, req.Email, enum.AccountRole(req.Role))
	if err != nil {
		return dto.TokenResponseDTO{}, err
	}

	refreshToken, err := util.GenerateRandomString(32)
	if err != nil {
		return dto.TokenResponseDTO{}, err
	}

	return dto.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    24 * 60 * 60,
	}, nil
}

func (s *authService) ChangePassword(ctx context.Context, req dto.ChangePasswordRequestDTO) error {
	userID, err := util.StringToUUID(req.UserID)
	if err != nil {
		return err
	}

	account, err := s.authRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	err = util.ComparePassword(account.Password, req.OldPassword)
	if err != nil {
		return constant.ErrInvalidCredentials
	}

	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	account.Password = hashedPassword
	return s.authRepo.Update(ctx, account)
}

func (s *authService) ResetPassword(ctx context.Context, req dto.ResetPasswordRequestDTO) error {
	account, err := s.authRepo.FindByEmail(ctx, req.ResetToken)
	if err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	account.Password = hashedPassword
	return s.authRepo.Update(ctx, account)
}
