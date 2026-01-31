package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"auth-service/internal/config"
	"auth-service/internal/controller"
	"auth-service/internal/repository"
	"auth-service/internal/service"

	authpb "auth-service/internal/pb/auth"

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
	accountRepository := repository.NewAccountRepository(db)
	refreshTokenRepository := repository.NewRefreshTokenRepository(db)
	authService := service.NewAuthService(accountRepository, refreshTokenRepository)
	authController := controller.NewAuthController(authService)

	// 🔹 Setup gRPC server
	grpcServer := grpc.NewServer()

	// 🔹 Register service to gRPC
	authpb.RegisterAuthServiceServer(grpcServer, authController)

	// 🔹 Running gRPC listener
	port := os.Getenv("AUTH_SERVICE_PORT")
	if port == "" {
		port = "10001"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("gagal listen di port %s: %v", port, err)
	}

	log.Printf("✅ Auth Service run on port %s 🚀", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
