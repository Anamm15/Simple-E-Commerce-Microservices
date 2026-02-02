package mapper

import (
	"notification-service/internal/model"
	notificationpb "notification-service/internal/pb/notification"
)

func MapToNotificationResponse(notification model.Notification) *notificationpb.Notification {
	return &notificationpb.Notification{
		Id:      notification.ID.String(),
		UserId:  notification.UserID.String(),
		Title:   notification.Title,
		Message: notification.Message,
		Type:    notification.Type,
		// CreatedAt: notification.CreatedAt,
	}
}

func MapToNotificationListResponse(notifications []model.Notification) []*notificationpb.Notification {
	var notificationList []*notificationpb.Notification
	for _, notification := range notifications {
		notificationList = append(notificationList, MapToNotificationResponse(notification))
	}
	return notificationList
}
