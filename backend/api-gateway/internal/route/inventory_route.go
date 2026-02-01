package route

import (
	"api-gateway/internal/controller"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterInventoryRoutes(router *gin.RouterGroup, inventoryController *controller.InventoryController) {
	inventoryGroup := router.Group("/inventory", middleware.Authenticate(), middleware.AuthorizeRole("ADMIN"))
	{
		inventoryGroup.PUT("/stock", inventoryController.UpdateStock)
	}
}
