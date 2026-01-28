package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"inventory-service/internal/config"
	"inventory-service/internal/controller"
	"inventory-service/internal/repository"
	"inventory-service/internal/service"

	inventory "inventory-service/internal/pb/inventory"

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
	inventoryRepository := repository.NewInventoryRepository(db)
	inventoryService := service.NewInventoryService(inventoryRepository)
	inventoryController := controller.NewInventoryController(inventoryService)

	// 🔹 Setup gRPC server
	grpcServer := grpc.NewServer()

	// 🔹 Register service to gRPC
	inventory.RegisterInventoryServiceServer(grpcServer, inventoryController)

	// 🔹 Running gRPC listener
	port := os.Getenv("INVENTORY_SERVICE_PORT")
	if port == "" {
		port = "10004"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	log.Printf("✅ Inventory Service run on port %s 🚀", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
