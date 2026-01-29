package grpc_client

import (
	"fmt"
	"log"

	"api-gateway/internal/pb/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentClient struct {
	Client paymentpb.PaymentServiceClient
}

func InitPaymentClient(host string, port string) (*PaymentClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to payment service: %v", err)
	}

	client := paymentpb.NewPaymentServiceClient(conn)

	log.Printf("✅ Connected to Payment Service at %s", target)

	return &PaymentClient{
		Client: client,
	}, nil
}
