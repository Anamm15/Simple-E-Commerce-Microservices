package route

import (
	"api-gateway/internal/controller"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authController *controller.AuthController) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/logout", authController.Logout)
		authGroup.POST("/refresh", authController.RefreshToken)
		authGroup.PATCH("/reset-password", authController.ResetPassword)
		protected := authGroup.Group("", middleware.Authenticate())
		{
			protected.PATCH("/change-password", authController.ChangePassword)
		}
	}
}
