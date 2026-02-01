package grpc_client

import (
	"fmt"
	"log"

	"api-gateway/internal/pb/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	Client authpb.AuthServiceClient
}

// InitAuthClient melakukan koneksi gRPC ke Service Auth
func InitAuthClient(host string, port string) (*AuthClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	// Menggunakan Insecure Credentials untuk development (tanpa TLS/SSL)
	// Untuk production, ganti dengan credentials.NewClientTLSFromFile(...)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Membuka koneksi
	conn, err := grpc.NewClient(target, opts...) // Atau grpc.Dial di versi lama
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %v", err)
	}

	// Create client dari kode generate protobuf
	client := authpb.NewAuthServiceClient(conn)

	log.Printf("✅ Connected to Auth Service at %s", target)

	return &AuthClient{
		Client: client,
	}, nil
}
