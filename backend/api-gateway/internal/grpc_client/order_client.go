package grpc_client

import (
	"fmt"
	"log"

	"api-gateway/internal/pb/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	Client orderpb.OrderServiceClient
}

func InitOrderClient(host string, port string) (*OrderClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to order service: %v", err)
	}

	client := orderpb.NewOrderServiceClient(conn)

	log.Printf("✅ Connected to Order Service at %s", target)

	return &OrderClient{
		Client: client,
	}, nil
}
