package dto

import (
	"user-service/internal/model"
	"user-service/pkg/util"
)

type AddressDetailResponseDTO struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Street    string `json:"street"`
	City      string `json:"city"`
	Province  string `json:"province"`
	ZipCode   string `json:"zip_code"`
	IsPrimary bool   `json:"is_primary"`
}

type AddAddressRequestDTO struct {
	UserID    string `json:"user_id" binding:"required"`
	Street    string `json:"street" binding:"required"`
	City      string `json:"city" binding:"required"`
	Province  string `json:"province" binding:"required"`
	ZipCode   string `json:"zip_code" binding:"required,numeric"`
	IsPrimary bool   `json:"is_primary"`
}

type UpdateAddressRequestDTO struct {
	AddressID string `json:"address_id" binding:"required"`
	UserID    string `json:"user_id" binding:"required"`
	Street    string `json:"street" binding:"required"`
	City      string `json:"city" binding:"required"`
	Province  string `json:"province" binding:"required"`
	ZipCode   string `json:"zip_code" binding:"required,numeric"`
	IsPrimary *bool  `json:"is_primary"`
}

func (r *AddAddressRequestDTO) ToModel() *model.Address {
	userID, err := util.StringToUUID(r.UserID)
	if err != nil {
		return nil
	}

	return &model.Address{
		UserID:    userID,
		Street:    r.Street,
		City:      r.City,
		Province:  r.Province,
		ZipCode:   r.ZipCode,
		IsPrimary: r.IsPrimary,
	}
}
