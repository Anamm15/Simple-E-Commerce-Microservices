package dto

import (
	"payment-service/internal/model"
	"payment-service/internal/util"
)

type CreatePaymentRequestDTO struct {
	OrderID  string `json:"order_id" binding:"required"`
	Amount   int64  `json:"amount" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type MidtransWebhookRequestDTO struct {
	OrderID           string `json:"order_id" binding:"required"`
	TransactionID     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status" binding:"required"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

func (dto *CreatePaymentRequestDTO) ToModel() model.Payment {
	OrderID, _ := util.StringToUUID(dto.OrderID)
	return model.Payment{
		OrderID: OrderID,
		Amount:  dto.Amount,
	}
}
