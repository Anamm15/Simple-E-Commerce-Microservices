package route

import (
	"api-gateway/internal/controller"

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
		authGroup.PATCH("/change-password", authController.ChangePassword)
	}
}
