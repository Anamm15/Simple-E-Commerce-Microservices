package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"user-service/internal/controller"
	"user-service/internal/infrastructure/dbms"
	"user-service/internal/repository"
	"user-service/internal/service"

	userpb "user-service/internal/pb/user"

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

	// 🔹 Initializing repository, service, and controller
	userRepository := repository.NewUserRepository(db)
	addressRepository := repository.NewAddressRepository(db)

	userService := service.NewUserService(userRepository)
	addressService := service.NewAddressService(addressRepository)

	userController := controller.NewUserController(userService, addressService)
	// 🔹 Setup gRPC server
	grpcServer := grpc.NewServer()

	// 🔹 Register service to gRPC
	userpb.RegisterUserServiceServer(grpcServer, userController)

	// 🔹 Running gRPC listener
	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "10002"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	log.Printf("✅ User Service run on port %s 🚀", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
