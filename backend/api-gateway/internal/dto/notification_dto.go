package dto

type GetNotificationsRequestDTO struct {
	Page  int32 `form:"page,default=1"`
	Limit int32 `form:"limit,default=20"`
}

type MarkNotificationReadRequestDTO struct{}

type SendNotificationRequestDTO struct {
	UserID  string `json:"user_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
	Type    string `json:"type" binding:"required"`
}
