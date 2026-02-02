package dto

import (
	"notification-service/internal/model"
	"notification-service/internal/util"
)

type GetNotificationsRequestDTO struct {
	UserID string `form:"user_id"`
	Page   int32  `form:"page,default=1"`
	Limit  int32  `form:"limit,default=20"`
	Sort   string `form:"sort"`
}

type SendNotificationRequestDTO struct {
	UserID  string `json:"user_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
	Type    string `json:"type" binding:"required"`
}

func (dto *SendNotificationRequestDTO) ToModel() *model.Notification {
	userID, _ := util.StringToUUID(dto.UserID)
	return &model.Notification{
		UserID:  userID,
		Title:   dto.Title,
		Message: dto.Message,
		Type:    dto.Type,
		IsRead:  false,
	}
}
