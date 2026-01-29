package dto

type CheckoutRequestDTO struct {
	UserID      string   `json:"user_id" binding:"required"`
	Address     string   `json:"address" binding:"required"`
	CourierCode string   `json:"courier_code" binding:"required"`
	ServiceCode string   `json:"service_code" binding:"required"`
	ProductIDs  []string `json:"product_ids" binding:"required"`
}

// 2. Get Order History (User)
type GetOrderHistoryRequestDTO struct {
	UserID       string `form:"user_id"`
	Page         int32  `form:"page,default=1"`
	Limit        int32  `form:"limit,default=10"`
	StatusFilter string `form:"status"`
	Sort         string `form:"sort"`
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
type UpdateOrderStatusRequestDTO struct {
	OrderID   string  `json:"order_id" binding:"required"`
	NewStatus string  `json:"new_status" binding:"required"`
	Notes     *string `json:"notes"`
}
