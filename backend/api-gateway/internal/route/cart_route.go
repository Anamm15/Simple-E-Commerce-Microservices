package route

import (
	"api-gateway/internal/controller"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCartRoutes(router *gin.RouterGroup, cartController *controller.CartController) {
	cartGroup := router.Group("/cart", middleware.Authenticate())
	{
		cartGroup.GET("", cartController.GetCart)
		cartGroup.POST("/items", cartController.AddItem)
		cartGroup.DELETE("/items/:product_id", cartController.RemoveItem)
		cartGroup.DELETE("/", cartController.ClearCart)
	}
}
