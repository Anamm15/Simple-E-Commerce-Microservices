package route

import (
	"api-gateway/internal/controller"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, userController *controller.UserController) {
	userGroup := router.Group("/users", middleware.Authenticate())
	{
		userGroup.GET("/profile", userController.GetUserProfile)
		userGroup.POST("", userController.CreateProfile)
		userGroup.PUT("/profile", userController.UpdateUserProfile)

		addressGroup := userGroup.Group("/addresses")
		{
			addressGroup.POST("/", userController.AddAddress)
			addressGroup.GET("/:id", userController.GetAddress)
			addressGroup.PUT("/:id", userController.UpdateAddress)
			addressGroup.DELETE("/:id", userController.DeleteAddress)
		}
	}
}
