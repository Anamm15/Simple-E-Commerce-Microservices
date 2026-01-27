package grpc_clients

import (
	"fmt"
	"log"

	"api-gateway/internal/pb/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ShippingClient struct {
	Client shippingpb.ShippingServiceClient
}

func InitShippingClient(host string, port string) (*ShippingClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to shipping service: %v", err)
	}

	client := shippingpb.NewShippingServiceClient(conn)

	log.Printf("✅ Connected to Shipping Service at %s", target)

	return &ShippingClient{
		Client: client,
	}, nil
}
