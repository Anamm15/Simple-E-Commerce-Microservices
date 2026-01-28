package dto

type CreatePaymentRequestDTO struct {
	OrderID string `json:"order_id" binding:"required"`
}

type MidtransWebhookRequestDTO struct {
	TransactionStatus string `json:"transaction_status" binding:"required"`
	OrderID           string `json:"order_id" binding:"required"`
	FraudStatus       string `json:"fraud_status"`
	PaymentType       string `json:"payment_type"`
	TransactionID     string `json:"transaction_id"`
}
