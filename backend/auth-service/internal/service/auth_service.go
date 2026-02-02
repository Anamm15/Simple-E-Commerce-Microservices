package service

import (
	"context"
	"errors"
	"time"

	"auth-service/internal/dto"
	"auth-service/internal/helper"
	"auth-service/internal/helper/constant"
	"auth-service/internal/model"
	authpb "auth-service/internal/pb/auth"
	"auth-service/internal/repository"
	"auth-service/pkg/util"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequestDTO) (dto.AccountResponseDTO, error)
	Login(ctx context.Context, req dto.LoginRequestDTO) (*authpb.LoginResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*authpb.TokenResponse, error)
	ChangePassword(ctx context.Context, req dto.ChangePasswordRequestDTO) error
	ResetPassword(ctx context.Context, req dto.ResetPasswordRequestDTO) error
}

type authService struct {
	authRepo         repository.AccountRepository
	refreshTokenRepo repository.RefreshTokenRepository
}

func NewAuthService(
	authRepo repository.AccountRepository,
	refreshTokenRepo repository.RefreshTokenRepository,
) AuthService {
	return &authService{
		authRepo:         authRepo,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequestDTO) (dto.AccountResponseDTO, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	newUserID := uuid.New()
	account := req.ToModel()
	account.UserID = newUserID
	account.Password = hashedPassword

	err = s.authRepo.Create(ctx, account)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	return *helper.MapAccountToAccountResponseDTO(account), nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequestDTO) (*authpb.LoginResponse, error) {
	account, err := s.authRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return &authpb.LoginResponse{}, err
	}

	err = util.ComparePassword(account.Password, req.Password)
	if err != nil {
		return &authpb.LoginResponse{}, constant.ErrInvalidCredentials
	}

	accessToken, err := util.GenerateJWT(account.UserID, account.Email, account.Role)
	if err != nil {
		return &authpb.LoginResponse{}, err
	}

	refreshTokenStr, err := util.GenerateRandomString(32)
	if err != nil {
		return &authpb.LoginResponse{}, err
	}

	familyID := uuid.New()

	refreshToken := model.RefreshToken{
		UserID:    account.UserID,
		Token:     refreshTokenStr,
		FamilyID:  familyID,
		IsUsed:    false,
		IsRevoked: false,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.refreshTokenRepo.Create(ctx, &refreshToken); err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		ExpiresIn:    24 * 60 * 60,
	}, nil
}

func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (*authpb.TokenResponse, error) {
	storedToken, err := s.refreshTokenRepo.GetByTokenHash(ctx, refreshToken)
	if err != nil {
		return &authpb.TokenResponse{}, errors.New("invalid refresh token")
	}

	// --- Reuse Detection ---
	if storedToken.IsUsed {
		_ = s.refreshTokenRepo.RevokeFamily(ctx, storedToken.FamilyID)
		return &authpb.TokenResponse{}, errors.New("refresh token reuse detected: security alert")
	}

	if storedToken.IsRevoked {
		return &authpb.TokenResponse{}, errors.New("token revoked")
	}
	if storedToken.ExpiresAt.Before(time.Now()) {
		return &authpb.TokenResponse{}, errors.New("token expired")
	}

	account, err := s.authRepo.FindByUserID(ctx, storedToken.UserID)
	if err != nil {
		return &authpb.TokenResponse{}, err
	}

	// --- Token Rotation ---
	newAccessToken, err := util.GenerateJWT(storedToken.UserID, account.Email, account.Role)
	if err != nil {
		return &authpb.TokenResponse{}, err
	}

	newRefreshTokenStr, err := util.GenerateRandomString(32)
	if err != nil {
		return &authpb.TokenResponse{}, err
	}

	newRefreshTokenRecord := model.RefreshToken{
		UserID:    storedToken.UserID,
		Token:     newRefreshTokenStr,
		FamilyID:  storedToken.FamilyID,
		IsUsed:    false,
		IsRevoked: false,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.refreshTokenRepo.RotateToken(ctx, storedToken.ID, &newRefreshTokenRecord)
	if err != nil {
		return &authpb.TokenResponse{}, err
	}

	response := &authpb.TokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshTokenStr,
		ExpiresIn:    7 * 24 * 60 * 60,
	}

	return response, nil
}

func (s *authService) ChangePassword(ctx context.Context, req dto.ChangePasswordRequestDTO) error {
	userID, err := util.StringToUUID(req.UserID)
	if err != nil {
		return err
	}

	account, err := s.authRepo.FindByUserID(ctx, userID)
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
