package grpc_clients

import (
	"fmt"
	"log"

	"api-gateway/internal/pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	Client userpb.UserServiceClient
}

// InitUserClient melakukan koneksi gRPC ke Service User
func InitUserClient(host string, port string) (*UserClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %v", err)
	}

	client := userpb.NewUserServiceClient(conn)

	log.Printf("✅ Connected to User Service at %s", target)

	return &UserClient{
		Client: client,
	}, nil
}
