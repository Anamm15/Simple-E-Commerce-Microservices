package route

import (
	"api-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterCartRoutes(router *gin.RouterGroup, cartController *controller.CartController) {
	cartGroup := router.Group("/cart")
	{
		cartGroup.GET("", cartController.GetCart)
		cartGroup.POST("/items", cartController.AddItem)
		cartGroup.DELETE("/items/:product_id", cartController.RemoveItem)
		cartGroup.DELETE("/", cartController.ClearCart)
	}
}
