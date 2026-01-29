package route

import (
	"api-gateway/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(router *gin.RouterGroup, orderController *controller.OrderController) {
	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("/checkout", orderController.Checkout)
		orderGroup.GET("/history", orderController.GetOrderHistory)
		orderGroup.GET("/:id", orderController.GetOrderDetail)
	}

	adminOrderGroup := router.Group("/admin/orders")
	{
		adminOrderGroup.GET("", orderController.GetAllOrders)
		adminOrderGroup.PUT("/:id/status", orderController.UpdateOrderStatus)
	}
}
