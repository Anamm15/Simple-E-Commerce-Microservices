package route

import (
	"api-gateway/internal/controller"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(router *gin.RouterGroup, orderController *controller.OrderController) {
	orderGroup := router.Group("/orders", middleware.Authenticate())
	{
		orderGroup.POST("/checkout", orderController.Checkout)
		orderGroup.GET("/history", orderController.GetOrderHistory)
		orderGroup.GET("/:id", orderController.GetOrderDetail)
	}

	adminOrderGroup := router.Group("/admin/orders", middleware.Authenticate(), middleware.AuthorizeRole("ADMIN"))
	{
		adminOrderGroup.GET("", orderController.GetAllOrders)
		adminOrderGroup.PATCH("/:id/status", orderController.UpdateOrderStatus)
	}
}
