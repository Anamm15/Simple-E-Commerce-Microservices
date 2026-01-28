package helper

import (
	"auth-service/internal/dto"
	"auth-service/internal/model"
)

func MapAccountToAccountResponseDTO(account *model.Account) *dto.AccountResponseDTO {
	return &dto.AccountResponseDTO{
		ID:       account.ID,
		Email:    account.Email,
		Username: account.Username,
		Role:     string(account.Role),
		Created:  account.CreatedAt,
	}
}
