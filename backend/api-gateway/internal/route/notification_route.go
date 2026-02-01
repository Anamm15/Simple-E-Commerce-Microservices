package route

import (
	"api-gateway/internal/controller"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(router *gin.RouterGroup, notifController *controller.NotificationController) {
	notifGroup := router.Group("/notifications", middleware.Authenticate())
	{
		notifGroup.GET("", notifController.GetNotifications)
		notifGroup.PUT("/:id/read", notifController.MarkAsRead)
	}

	adminGroup := router.Group("/admin/notifications", middleware.Authenticate(), middleware.AuthorizeRole("ADMIN"))
	{
		adminGroup.POST("/send", notifController.SendNotification)
	}
}
