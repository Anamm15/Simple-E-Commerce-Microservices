package repository

import (
	"context"
	"strings"

	"notification-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	GetUserNotifications(ctx context.Context, userID uuid.UUID, limit *int32, offset *int32, sort *string) ([]model.Notification, error)
	GetUnreadNotificationsCount(ctx context.Context, userID uuid.UUID) (int64, error)
	Create(ctx context.Context, notification *model.Notification) error
	MarkAsRead(ctx context.Context, notificationID uuid.UUID, userID uuid.UUID) error
	Delete(ctx context.Context, notificationID uuid.UUID, userID uuid.UUID) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) GetUserNotifications(
	ctx context.Context,
	userID uuid.UUID,
	limit *int32,
	offset *int32,
	sort *string,
) ([]model.Notification, error) {
	var notifications []model.Notification

	query := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID)

	if sort != nil {
		switch strings.ToLower(*sort) {
		case "oldest":
			query = query.Order("created_at ASC")
		default:
			query = query.Order("created_at DESC")
		}
	} else {
		query = query.Order("created_at DESC")
	}

	if limit != nil {
		query = query.Limit(int(*limit))
	}
	if offset != nil {
		query = query.Offset(int(*offset))
	}

	err := query.Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *notificationRepository) GetUnreadNotificationsCount(ctx context.Context, userID uuid.UUID) (int64, error) {
	var notificationCount int64

	err := r.db.
		WithContext(ctx).
		Where("user_id = ? AND is_read = ?", userID, false).
		Count(&notificationCount).
		Error
	if err != nil {
		return 0, err
	}

	return notificationCount, nil
}

func (r *notificationRepository) Create(
	ctx context.Context,
	notification *model.Notification,
) error {
	return r.db.
		WithContext(ctx).
		Create(notification).
		Error
}

func (r *notificationRepository) MarkAsRead(
	ctx context.Context,
	notificationID uuid.UUID,
	userID uuid.UUID,
) error {
	result := r.db.
		WithContext(ctx).
		Model(&model.Notification{}).
		Where("id = ? AND user_id = ?", notificationID, userID).
		Update("is_read", true)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (r *notificationRepository) Delete(
	ctx context.Context,
	notificationID uuid.UUID,
	userID uuid.UUID,
) error {
	result := r.db.
		WithContext(ctx).
		Delete(
			&model.Notification{},
			"id = ? AND user_id = ?",
			notificationID,
			userID,
		)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
