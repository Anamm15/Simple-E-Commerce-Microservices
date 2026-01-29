package dto

type UpdateUserProfileRequestDTO struct {
	FullName    string `json:"full_name" binding:"required,min=2"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
}

type AddAddressRequestDTO struct {
	Street    string `json:"street" binding:"required"`
	City      string `json:"city" binding:"required"`
	Province  string `json:"province" binding:"required"`
	ZipCode   string `json:"zip_code" binding:"required,numeric"`
	IsPrimary bool   `json:"is_primary"`
}

type UpdateAddressRequestDTO struct {
	Street    string `json:"street" binding:"required"`
	City      string `json:"city" binding:"required"`
	Province  string `json:"province" binding:"required"`
	ZipCode   string `json:"zip_code" binding:"required,numeric"`
	IsPrimary bool   `json:"is_primary"`
}
