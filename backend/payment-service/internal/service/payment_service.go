package service

import (
	"context"
	"fmt"

	"payment-service/internal/dto"
	"payment-service/internal/helper/mapper"
	"payment-service/internal/infrastructure/kafka"
	paymentpb "payment-service/internal/pb/payment"
	"payment-service/internal/repository"
	"payment-service/pkg/util"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, request dto.CreatePaymentRequestDTO) (*paymentpb.PaymentResponse, error)
	WebhookPayment(ctx context.Context, request dto.MidtransWebhookRequestDTO) error
}

type paymentService struct {
	paymentRepository repository.PaymentRepository
	producer          kafka.Producer
	snapClient        snap.Client
}

func NewPaymentService(
	paymentRepository repository.PaymentRepository,
	producer kafka.Producer,
	midtransServerKey string,
) PaymentService {
	var client snap.Client
	client.New(midtransServerKey, midtrans.Sandbox)
	return &paymentService{
		paymentRepository: paymentRepository,
		producer:          producer,
		snapClient:        client,
	}
}

func (s *paymentService) CreatePayment(ctx context.Context, request dto.CreatePaymentRequestDTO) (*paymentpb.PaymentResponse, error) {
	payment := request.ToModel()

	if err := s.paymentRepository.Create(ctx, &payment); err != nil {
		return &paymentpb.PaymentResponse{}, err
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.OrderID,
			GrossAmt: int64(request.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: request.FullName,
			Email: request.Email,
		},
	}

	snapResp, err := s.snapClient.CreateTransaction(req)
	if err != nil {
		return &paymentpb.PaymentResponse{}, err
	}

	paymentGatewayResponse := &paymentpb.PaymentResponse{
		RedirectUrl: snapResp.RedirectURL,
		Token:       snapResp.Token,
	}

	return paymentGatewayResponse, nil
}

func (s *paymentService) WebhookPayment(ctx context.Context, request dto.MidtransWebhookRequestDTO) error {
	orderID, err := util.StringToUUID(request.OrderID)
	if err != nil {
		return fmt.Errorf("invalid order id: %w", err)
	}

	paymentStatus, _ := mapper.MapMidtransStatus(request.TransactionStatus, request.FraudStatus)

	payment, err := s.paymentRepository.FindByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	payment.Status = string(paymentStatus)
	payment.MidtransTransactionID = &request.TransactionID
	payment.PaymentMethod = &request.PaymentType

	return nil
}
