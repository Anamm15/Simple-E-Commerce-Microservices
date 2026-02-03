package dto

type CreateUserProfileRequestDTO struct {
	FullName    string `json:"full_name" binding:"required,min=2"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
}

type UpdateUserProfileRequestDTO struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type AddAddressRequestDTO struct {
	Street    string `json:"street" binding:"required"`
	City      string `json:"city" binding:"required"`
	Province  string `json:"province" binding:"required"`
	ZipCode   string `json:"zip_code" binding:"required"`
	IsPrimary bool   `json:"is_primary"`
}

type UpdateAddressRequestDTO struct {
	Street    string `json:"street" `
	City      string `json:"city" `
	Province  string `json:"province" `
	ZipCode   string `json:"zip_code"`
	IsPrimary bool   `json:"is_primary"`
}
