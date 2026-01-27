package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	grpc_clients "api-gateway/internal/grpc_clients"

	"api-gateway/internal/routes"
)

func main() {
	// 1. Load Environment Variables (.env)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: No .env file found, using system environment variables")
	}

	// 2. Initialize all gRPC Clients
	log.Println("🔌 Initializing gRPC Clients...")

	// Auth Service
	authSvc, err := grpc_clients.InitAuthClient(os.Getenv("AUTH_SERVICE_HOST"), os.Getenv("AUTH_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to init Auth Client: %v", err)
	}

	// User Service
	userSvc, err := grpc_clients.InitUserClient(os.Getenv("USER_SERVICE_HOST"), os.Getenv("USER_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to init User Client: %v", err)
	}

	// Product Service
	productSvc, err := grpc_clients.InitProductClient(os.Getenv("PRODUCT_SERVICE_HOST"), os.Getenv("PRODUCT_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to init Product Client: %v", err)
	}

	// Inventory Service
	inventorySvc, err := grpc_clients.InitInventoryClient(os.Getenv("INVENTORY_SERVICE_HOST"), os.Getenv("INVENTORY_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to init Inventory Client: %v", err)
	}

	// Cart Service
	cartSvc, err := grpc_clients.InitCartClient(os.Getenv("CART_SERVICE_HOST"), os.Getenv("CART_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to init Cart Client: %v", err)
	}

	// Order Service
	orderSvc, err := grpc_clients.InitOrderClient(os.Getenv("ORDER_SERVICE_HOST"), os.Getenv("ORDER_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to init Order Client: %v", err)
	}

	// Payment Service
	paymentSvc, err := grpc_clients.InitPaymentClient(os.Getenv("PAYMENT_SERVICE_HOST"), os.Getenv("PAYMENT_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to init Payment Client: %v", err)
	}

	// Shipping Service
	shippingSvc, err := grpc_clients.InitShippingClient(os.Getenv("SHIPPING_SERVICE_HOST"), os.Getenv("SHIPPING_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to init Shipping Client: %v", err)
	}

	// Notification Service
	notifSvc, err := grpc_clients.InitNotificationClient(os.Getenv("NOTIFICATION_SERVICE_HOST"), os.Getenv("NOTIFICATION_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to init Notification Client: %v", err)
	}

	// 3. Setup and Register Routes
	log.Println("🛣️  Registering Routes...")
	r := routes.SetupRouter(
		authSvc.Client,
		userSvc.Client,
		productSvc.Client,
		inventorySvc.Client,
		cartSvc.Client,
		orderSvc.Client,
		paymentSvc.Client,
		shippingSvc.Client,
		notifSvc.Client,
	)

	// Health Check Endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "UP",
			"service": "API Gateway",
		})
	})

	// 4. Run Server
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	log.Printf("🚀 API Gateway running on port :%s", appPort)
	if err := r.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
