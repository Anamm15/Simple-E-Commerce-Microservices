package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"product-service/internal/config"
	"product-service/internal/controller"
	"product-service/internal/grpc_client"
	"product-service/internal/repository"
	"product-service/internal/service"

	productpb "product-service/internal/pb/product"

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

	//🔹 Connecting to gRPC
	productClient, err := grpc_client.InitInventoryClient(os.Getenv("INVENTORY_SERVICE_HOST"), os.Getenv("INVENTORY_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to connect to product service: %v", err)
	}

	// 🔹 Initializing repository, service, and controller
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(cartRepository, productClient.Client)
	productController := controller.NewProductController(cartService)

	// 🔹 Setup gRPC server
	grpcServer := grpc.NewServer()

	// 🔹 Register service to gRPC
	productpb.RegisterProductServiceServer(grpcServer, productController)

	// 🔹 Running gRPC listener
	port := os.Getenv("PRODUCT_SERVICE_PORT")
	if port == "" {
		port = "10003"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	log.Printf("✅ Cart Service run on port %s 🚀", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
