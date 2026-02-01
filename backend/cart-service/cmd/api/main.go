package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"cart-service/internal/config"
	"cart-service/internal/controller"
	"cart-service/internal/grpc_client"
	"cart-service/internal/infrastructure/dbms"
	"cart-service/internal/infrastructure/kafka"
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

	// 🔹 Initialize Root Context for Graceful Shutdown
	// This context manages the lifecycle of background processes (like the Kafka Consumer).
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 🔹 Connecting to database
	db := dbms.ConnectDatabase()

	// 🔹 Load kafka configuration
	kafkaConfig := config.LoadKafkaConfig()

	// 🔹 Initializing Kafka consumer
	consumer := kafka.NewConsumer(kafkaConfig.KafkaBrokers)

	// 🔹 Spawn Kafka Consumer (Asynchronous Worker)
	// Running in a separate goroutine to avoid blocking the main thread.
	go func() {
		log.Println("🎧 Kafka Consumer Worker started...")

		// Define the MessageHandler closure
		messageHandler := func(ctx context.Context, key []byte, value []byte) error {
			log.Printf("📥 Inbound Message - Key: %s, Payload Size: %d bytes", string(key), len(value))

			// TODO: Invoke service layer to process the payload
			// return paymentService.ProcessEvent(ctx, value)
			return nil
		}

		// Start consuming from the 'orders' topic within the 'payment-service-group'
		// This block blocks until the context is canceled or an error occurs.
		err := consumer.StartConsumerGroup(ctx, "payment-service-group", []string{"orders"}, messageHandler)
		if err != nil {
			log.Printf("❌ Consumer Group terminated abruptly: %v", err)
		}
	}()

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
