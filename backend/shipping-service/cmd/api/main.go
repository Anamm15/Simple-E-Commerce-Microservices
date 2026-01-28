package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"shipping-service/internal/config"
	"shipping-service/internal/controller"
	"shipping-service/internal/repository"
	"shipping-service/internal/service"

	shippingpb "shipping-service/internal/pb/shipping"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// checking environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := fmt.Sprintf(".env.%s", env)
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠️ Cannot load %s, using default .env", envFile)
		_ = godotenv.Load(".env")
	}

	// 🔹 Connecting to database
	db := config.ConnectDatabase()

	// 🔹 Initializing repository, service, and controller
	shippingRepository := repository.NewShippingRepository(db)
	shippingService := service.NewShippingService(shippingRepository)
	shippingController := controller.NewShippingController(shippingService)
	// 🔹 Setup gRPC server
	grpcServer := grpc.NewServer()

	// 🔹 Register service to gRPC
	shippingpb.RegisterShippingServiceServer(grpcServer, shippingController)

	// 🔹 Running gRPC listener
	port := os.Getenv("SHIPPING_SERVICE_PORT")
	if port == "" {
		port = "10006"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	log.Printf("✅ Shipping Service run on port %s 🚀", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
