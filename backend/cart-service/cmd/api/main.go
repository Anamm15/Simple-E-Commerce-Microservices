package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"cart-service/internal/controller"
	"cart-service/internal/grpc_client"
	"cart-service/internal/infrastructure/dbms"
	"cart-service/internal/repository"
	"cart-service/internal/service"

	cartpb "cart-service/internal/pb/cart"

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
	db := dbms.ConnectDatabase()

	//🔹 Connecting to gRPC
	productClient, err := grpc_client.InitProductClient(os.Getenv("PRODUCT_SERVICE_HOST"), os.Getenv("PRODUCT_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to connect to product service: %v", err)
	}

	// 🔹 Initializing repository, service, and controller
	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository, productClient.Client)
	cartController := controller.NewCartController(cartService)

	// 🔹 Setup gRPC server
	grpcServer := grpc.NewServer()

	// 🔹 Register service to gRPC
	cartpb.RegisterCartServiceServer(grpcServer, cartController)

	// 🔹 Running gRPC listener
	port := os.Getenv("CART_SERVICE_PORT")
	if port == "" {
		port = "10005"
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
