package grpc_client

import (
	"fmt"
	"log"

	productpb "cart-service/internal/pb/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	Client productpb.ProductServiceClient
}

func InitProductClient(host string, port string) (*ProductClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to product service: %v", err)
	}

	client := productpb.NewProductServiceClient(conn)

	log.Printf("✅ Connected to Product Service at %s", target)

	return &ProductClient{
		Client: client,
	}, nil
}
