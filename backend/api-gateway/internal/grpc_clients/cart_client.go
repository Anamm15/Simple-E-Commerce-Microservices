package grpc_clients

import (
	"fmt"
	"log"

	"api-gateway/internal/pb/cart"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CartClient struct {
	Client cartpb.CartServiceClient
}

func InitCartClient(host string, port string) (*CartClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to cart service: %v", err)
	}

	client := cartpb.NewCartServiceClient(conn)

	log.Printf("✅ Connected to Cart Service at %s", target)

	return &CartClient{
		Client: client,
	}, nil
}
