package dto

import "user-service/internal/model"

type CreateProfileRequestDTO struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateUserProfileRequestDTO struct {
	UserID      string `json:"user_id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

func (dto *CreateProfileRequestDTO) ToModel() *model.User {
	return &model.User{
		FullName:    dto.FullName,
		PhoneNumber: dto.PhoneNumber,
	}
}
