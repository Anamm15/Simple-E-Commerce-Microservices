package grpc_client

import (
	"fmt"
	"log"

	inventorypb "order-service/internal/pb/inventory"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InventoryClient struct {
	Client inventorypb.InventoryServiceClient
}

func InitInventoryClient(host string, port string) (*InventoryClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to inventory service: %v", err)
	}

	client := inventorypb.NewInventoryServiceClient(conn)

	log.Printf("✅ Connected to Inventory Service at %s", target)

	return &InventoryClient{
		Client: client,
	}, nil
}
