package route

import (
	"api-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterShippingRoutes(router *gin.RouterGroup, shippingController *controller.ShippingController) {
	paymentGroup := router.Group("/shipments")
	{
		paymentGroup.POST("/offer", shippingController.OfferShippingCost)
	}
}
