package grpc_client

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	notificationpb "order-service/internal/pb/notification"
)

type NotificationClient struct {
	Client notificationpb.NotificationServiceClient
}

func InitNotificationClient(host string, port string) (*NotificationClient, error) {
	target := fmt.Sprintf("%s:%s", host, port)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to notification service: %v", err)
	}

	client := notificationpb.NewNotificationServiceClient(conn)

	log.Printf("✅ Connected to Notification Service at %s", target)

	return &NotificationClient{
		Client: client,
	}, nil
}
