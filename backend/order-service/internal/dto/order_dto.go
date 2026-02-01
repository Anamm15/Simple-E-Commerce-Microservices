package dto

import "order-service/internal/helper/enum"

type AddressRequestDTO struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	ZipCode  string `json:"zip_code"`
	Province string `json:"province"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

type ProductCheckout struct {
	ProductID string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

type CheckoutRequestDTO struct {
	UserID      string            `json:"user_id"`
	Email       string            `json:"email"`
	Address     AddressRequestDTO `json:"address"`
	CourierCode string            `json:"courier_code"`
	ServiceCode string            `json:"service_code"`
	Products    []ProductCheckout `json:"products"`
}

// 2. Get Order History (User)
type GetOrderHistoryRequestDTO struct {
	UserID       string           `form:"user_id"`
	Page         int32            `form:"page,default=1"`
	Limit        int32            `form:"limit,default=10"`
	StatusFilter enum.OrderStatus `form:"status"`
	Sort         string           `form:"sort"`
}

type AdminOrderFilterRequestDTO struct {
	Page      int32            `form:"page,default=1"`
	Limit     int32            `form:"limit,default=20"`
	UserID    string           `form:"user_id"`
	Status    enum.OrderStatus `form:"status"`
	DateStart string           `form:"date_start"`
	DateEnd   string           `form:"date_end"`
}

// 5. Update Order Status (Admin)
type UpdateOrderStatusRequestDTO struct {
	OrderID   string           `json:"order_id" binding:"required"`
	NewStatus enum.OrderStatus `json:"new_status" binding:"required"`
	Notes     *string          `json:"notes"`
}
