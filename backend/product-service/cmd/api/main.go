package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"product-service/internal/config"
	"product-service/internal/controller"
	"product-service/internal/grpc_client"
	"product-service/internal/infrastructure/cloud"
	"product-service/internal/infrastructure/dbms"
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
	db := dbms.ConnectDatabase()

	// 🔹 Load Cloudinary Config
	cloudConfig := config.LoadCloudinaryConfig()

	// 🔹 Initialize Cloudinary Adapter
	cloudSvc, err := cloud.NewCloudinaryService(
		cloudConfig.CloudName,
		cloudConfig.APIKey,
		cloudConfig.APISecret,
		cloudConfig.UploadFolder,
	)
	if err != nil {
		log.Fatalf("FATAL: Failed to init Cloudinary: %v", err)
	}

	// 🔹 Connecting to inventory
	inventorySvc, err := grpc_client.InitInventoryClient(os.Getenv("INVENTORY_SERVICE_HOST"), os.Getenv("INVENTORY_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to connect to cart service: %v", err)
	}

	// 🔹 Initializing repository, service, and controller
	productRepository := repository.NewProductRepository(db)
	reviewRepository := repository.NewReviewRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	ImageRepository := repository.NewImageRepository(db)
	productService := service.NewProductService(productRepository, reviewRepository, categoryRepository, inventorySvc.Client, cloudSvc)
	imageService := service.NewImageService(ImageRepository, cloudSvc)
	productController := controller.NewProductController(productService, imageService)

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

	log.Printf("✅ Product Service run on port %s 🚀", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
