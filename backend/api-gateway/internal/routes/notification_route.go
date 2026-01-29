package routes

import (
	"api-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(router *gin.RouterGroup, notifController *controller.NotificationController) {
	notifGroup := router.Group("/notifications")
	{
		notifGroup.GET("", notifController.GetNotifications)
		notifGroup.PUT("/:id/read", notifController.MarkAsRead)
	}

	adminGroup := router.Group("/admin/notifications")
	{
		adminGroup.POST("/send", notifController.SendNotification)
	}
}
