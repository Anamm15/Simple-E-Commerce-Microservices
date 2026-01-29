package controller

import (
	"context"

	"payment-service/internal/dto"
	paymentpb "payment-service/internal/pb/payment"
	"payment-service/internal/service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type paymentController struct {
	paymentpb.UnimplementedPaymentServiceServer
	paymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) paymentpb.PaymentServiceServer {
	return &paymentController{paymentService: paymentService}
}

func (c *paymentController) CreatePayment(ctx context.Context, req *paymentpb.CreatePaymentRequest) (*paymentpb.PaymentResponse, error) {
	input := dto.CreatePaymentRequestDTO{
		OrderID:  req.OrderId,
		Amount:   req.Amount,
		Email:    req.CustomerEmail,
		FullName: req.CustomerName,
	}
	createdPayment, err := c.paymentService.CreatePayment(ctx, input)
	if err != nil {
		return nil, err
	}

	return createdPayment, nil
}

func (c *paymentController) WebhookPayment(ctx context.Context, req *paymentpb.WebhookRequest) (*emptypb.Empty, error) {
	input := dto.MidtransWebhookRequestDTO{
		OrderID:           req.OrderId,
		TransactionID:     req.TransactionId,
		TransactionStatus: req.TransactionStatus,
		PaymentType:       req.PaymentType,
		FraudStatus:       req.FraudStatus,
	}

	if err := c.paymentService.WebhookPayment(ctx, input); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
