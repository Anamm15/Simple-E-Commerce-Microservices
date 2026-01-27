package dto

type CheckoutRequestDTO struct {
	AddressID     string `json:"address_id" binding:"required"`
	CourierCode   string `json:"courier_code" binding:"required"`
	ServiceCode   string `json:"service_code" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

// 2. Get Order History (User)
// Endpoint: GET /orders/history
type GetOrderHistoryRequestDTO struct {
	Page         int32  `form:"page,default=1"`
	Limit        int32  `form:"limit,default=10"`
	StatusFilter string `form:"status"`
}

type AdminOrderFilterRequestDTO struct {
	Page      int32  `form:"page,default=1"`
	Limit     int32  `form:"limit,default=20"`
	UserID    string `form:"user_id"`
	Status    string `form:"status"`
	DateStart string `form:"date_start"`
	DateEnd   string `form:"date_end"`
}

// 5. Update Order Status (Admin)
// Endpoint: PUT /admin/orders/:order_id/status
type UpdateOrderStatusRequestDTO struct {
	NewStatus string `json:"new_status" binding:"required"`
	Notes     string `json:"notes"`
}
