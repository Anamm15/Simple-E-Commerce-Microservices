package grpc_clients

import (
	"fmt"
	"log"

	"api-gateway/internal/pb/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	Client productpb.ProductServiceClient
}

func InitProductClient(host string, port string) (*ProductClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	// Opsi koneksi (Insecure untuk dev environment)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// Opsional: Konfigurasi Load Balancing jika ada banyak instance service product
		// grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
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
