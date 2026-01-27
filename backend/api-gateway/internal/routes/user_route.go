package routes

import (
	"api-gateway/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, userController *controller.UserController) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/profile", userController.GetUserProfile)
		userGroup.PUT("/profile", userController.UpdateUserProfile)

		addressGroup := userGroup.Group("/addresses")
		{
			addressGroup.POST("/", userController.AddAddress)
			addressGroup.GET("/:address_id", userController.GetAddress)
			addressGroup.PUT("/:address_id", userController.UpdateAddress)
			addressGroup.DELETE("/:address_id", userController.DeleteAddress)
		}
	}
}
