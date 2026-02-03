package service

import (
	"context"

	"notification-service/internal/dto"
	"notification-service/internal/helper"
	"notification-service/internal/helper/constant"
	"notification-service/internal/helper/mapper"
	notificationpb "notification-service/internal/pb/notification"
	"notification-service/internal/repository"
	"notification-service/internal/util"
)

type NotificationService interface {
	GetUserNotifications(ctx context.Context, userID string, page *int32, offset *int32, sort *string) (*notificationpb.NotificationList, error)
	CreateNotification(ctx context.Context, notification dto.SendNotificationRequestDTO) error
	MarkAsRead(ctx context.Context, notificationID string, userID string) error
}

type notificationService struct {
	notificationRepo repository.NotificationRepository
}

func NewNotificationService(notificationRepo repository.NotificationRepository) NotificationService {
	return &notificationService{notificationRepo: notificationRepo}
}

func (s *notificationService) GetUserNotifications(ctx context.Context, userID string, page *int32, limit *int32, sort *string) (*notificationpb.NotificationList, error) {
	userIDParsed, err := util.StringToUUID(userID)
	if err != nil {
		return nil, err
	}

	var limitParsed, pageParsed int32
	var sortParsed string

	if limit == nil {
		limitParsed = constant.DefaultLimit
	} else {
		limitParsed = *limit
	}

	if page == nil {
		pageParsed = constant.DefaultPage
	} else {
		pageParsed = *page
	}

	if sort == nil {
		sortParsed = constant.DefaultSort
	} else {
		sortParsed = *sort
	}

	var offset int32
	if page != nil {
		offset = helper.CalculateOffset(pageParsed, limitParsed)
	}
	notifications, err := s.notificationRepo.GetUserNotifications(ctx, userIDParsed, limitParsed, offset, sortParsed)
	if err != nil {
		return nil, err
	}

	notificationCount, err := s.notificationRepo.GetUnreadNotificationsCount(ctx, userIDParsed)
	if err != nil {
		return nil, err
	}

	return &notificationpb.NotificationList{
		Notifications: mapper.MapToNotificationListResponse(notifications),
		UnreadCount:   notificationCount,
	}, nil
}

func (s *notificationService) CreateNotification(ctx context.Context, notification dto.SendNotificationRequestDTO) error {
	notificationModel := notification.ToModel()
	return s.notificationRepo.Create(ctx, notificationModel)
}

func (s *notificationService) MarkAsRead(ctx context.Context, notificationID string, userID string) error {
	notificationIDParsed, err := util.StringToUUID(notificationID)
	if err != nil {
		return err
	}
	userIDParsed, err := util.StringToUUID(userID)
	if err != nil {
		return err
	}
	return s.notificationRepo.MarkAsRead(ctx, notificationIDParsed, userIDParsed)
}
