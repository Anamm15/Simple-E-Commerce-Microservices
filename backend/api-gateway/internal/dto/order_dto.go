package dto

type ProductCheckoutRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required,min=1"`
}

type AddressOrderSnapshot struct {
	Street        string `json:"street" binding:"required"`
	City          string `json:"city" binding:"required"`
	Province      string `json:"province" binding:"required"`
	ZipCode       string `json:"zip_code" binding:"required,numeric"`
	ReceiverName  string `json:"receiver_name" binding:"required"`
	ReceiverPhone string `json:"receiver_phone" binding:"required,e164"`
}

type CheckoutRequestDTO struct {
	Address     AddressOrderSnapshot     `json:"address" binding:"required"`
	CourierCode string                   `json:"courier_code" binding:"required"`
	ServiceCode string                   `json:"service_code" binding:"required"`
	Products    []ProductCheckoutRequest `json:"products" binding:"required"`
}

// 2. Get Order History (User)
// Endpoint: GET /orders/history
type GetOrderHistoryRequestDTO struct {
	Page         int32  `json:"page"`
	Limit        int32  `json:"limit"`
	StatusFilter string `json:"status"`
	Sort         string `json:"sort"`
}

type AdminOrderFilterRequestDTO struct {
	Page      int32  `json:"page"`
	Limit     int32  `json:"limit"`
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
	DateStart string `json:"date_start"`
	DateEnd   string `json:"date_end"`
}

// 5. Update Order Status (Admin)
// Endpoint: PUT /admin/orders/:order_id/status
type UpdateOrderStatusRequestDTO struct {
	NewStatus string `json:"new_status" binding:"required"`
	Notes     string `json:"notes"`
}
