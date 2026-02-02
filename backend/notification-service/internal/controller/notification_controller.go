package controller

import (
	"context"

	"notification-service/internal/dto"
	notificationpb "notification-service/internal/pb/notification"
	"notification-service/internal/service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type notificationController struct {
	notificationpb.UnimplementedNotificationServiceServer
	notificationService service.NotificationService
}

func NewNotificationController(notificationService service.NotificationService) *notificationController {
	return &notificationController{
		notificationService: notificationService,
	}
}

func (c *notificationController) GetNotifications(ctx context.Context, request *notificationpb.GetNotificationsRequest) (*notificationpb.NotificationList, error) {
	return c.notificationService.GetUserNotifications(ctx, request.UserId, &request.Page, &request.Limit, &request.Sort)
}

func (c *notificationController) MarkAsRead(ctx context.Context, request *notificationpb.MarkAsReadRequest) (*emptypb.Empty, error) {
	if err := c.notificationService.MarkAsRead(ctx, request.NotificationId, request.UserId); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (c *notificationController) SendNotification(ctx context.Context, request *notificationpb.SendNotificationRequest) (*emptypb.Empty, error) {
	input := dto.SendNotificationRequestDTO{
		UserID:  request.UserId,
		Title:   request.Title,
		Type:    request.Type,
		Message: request.Message,
	}

	if err := c.notificationService.CreateNotification(ctx, input); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
