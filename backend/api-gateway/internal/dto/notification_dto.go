package dto

type GetNotificationsRequestDTO struct {
	Page  int32  `json:"page"`
	Limit int32  `json:"limit"`
	Sort  string `json:"sort"`
}

type SendNotificationRequestDTO struct {
	UserID  string `json:"user_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
	Type    string `json:"type" binding:"required"`
}
