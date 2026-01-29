package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"payment-service/internal/config"
	"payment-service/internal/controller"
	"payment-service/internal/infrastructure/dbms"
	"payment-service/internal/infrastructure/kafka"
	"payment-service/internal/repository"
	"payment-service/internal/service"

	paymentpb "payment-service/internal/pb/payment"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// checking environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	midtransServerKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if midtransServerKey == "" {
		log.Fatal("⚠️ MIDTRANS_SERVER_KEY is not set")
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

	// 🔹 Initializing repository, service, and controller
	paymentRepository := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepository, producer, midtransServerKey)
	paymentController := controller.NewPaymentController(paymentService)

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

	// 🔹 Setup gRPC server
	grpcServer := grpc.NewServer()

	// 🔹 Register service to gRPC
	paymentpb.RegisterPaymentServiceServer(grpcServer, paymentController)

	// 🔹 Running gRPC listener
	port := os.Getenv("PAYMENT_SERVICE_PORT")
	if port == "" {
		port = "10007"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	log.Printf("✅ Payment Service run on port %s 🚀", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// 🔹 Start gRPC server
	go func() {
		log.Printf("✅ gRPC Server listening on port %s 🚀", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("FATAL: Failed to serve gRPC: %v", err)
		}
	}()

	// 🔹 Graceful Shutdown Mechanism
	// Block main thread until an OS signal (SIGINT/SIGTERM) is received.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit // Block here
	log.Println("⚠️  Shutdown signal received. Initiating termination sequence...")

	// Step A: Propagate cancellation context to stop the Kafka Consumer loop
	cancel()

	// Step B: Gracefully stop the gRPC server (drains active connections)
	grpcServer.GracefulStop()

	// Step C: Allow a brief window for cleanup operations
	time.Sleep(1 * time.Second)
	log.Println("👋 Application exited successfully.")
}
