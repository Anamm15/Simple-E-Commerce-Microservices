package route

import (
	"api-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(router *gin.RouterGroup, paymentController *controller.PaymentController) {
	paymentGroup := router.Group("/payments")
	{
		paymentGroup.POST("/webhook", paymentController.WebhookPayment)
	}
}
