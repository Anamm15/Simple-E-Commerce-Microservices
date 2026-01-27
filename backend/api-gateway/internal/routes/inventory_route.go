package routes

import (
	"api-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterInventoryRoutes(router *gin.RouterGroup, inventoryController *controller.InventoryController) {
	inventoryGroup := router.Group("/inventory")
	{
		inventoryGroup.PUT("/stock", inventoryController.UpdateStock)
	}
}
