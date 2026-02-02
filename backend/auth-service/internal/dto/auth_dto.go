package dto

import (
	"time"

	"auth-service/internal/model"

	"github.com/google/uuid"
)

type AccountResponseDTO struct {
	ID       uuid.UUID
	Email    string
	Username string
	Role     string
	Created  time.Time
}

type RegisterRequestDTO struct {
	Email    string
	Username string
	Password string
}

type LoginRequestDTO struct {
	Email    string
	Password string
}

type ChangePasswordRequestDTO struct {
	UserID      string
	OldPassword string
	NewPassword string
}

type ResetPasswordRequestDTO struct {
	ResetToken  string
	NewPassword string
}

type Claims struct {
	UserID string
	Role   string
	Email  string
}

func (dto *RegisterRequestDTO) ToModel() *model.Account {
	return &model.Account{
		Email:    dto.Email,
		Username: dto.Username,
	}
}
