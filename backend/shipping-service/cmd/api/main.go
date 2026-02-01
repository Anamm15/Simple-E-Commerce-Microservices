package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"shipping-service/internal/config"
	"shipping-service/internal/controller"
	"shipping-service/internal/infrastructure/dbms"
	"shipping-service/internal/infrastructure/kafka"
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

	// 🔹 Initialize Root Context for Graceful Shutdown
	// This context manages the lifecycle of background processes (like the Kafka Consumer).
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 🔹 Connecting to database
	db := dbms.ConnectDatabase()

	// 🔹 Load kafka configuration
	kafkaConfig := config.LoadKafkaConfig()

	// 🔹 Initializing Kafka producer
	producer, err := kafka.NewProducer(kafkaConfig.KafkaBrokers)
	if err != nil {
		log.Fatalf("Failed to initialize kafka producer: %v", err)
	}
	// Ensure producer is closed upon application exit
	defer func() {
		log.Println("Closing Kafka Producer connection...")
		producer.Close()
	}()

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
