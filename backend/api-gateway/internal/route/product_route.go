package route

import (
	"api-gateway/internal/controller"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.RouterGroup, productController *controller.ProductController) {
	productGroup := router.Group("/products")
	{
		productGroup.GET("", productController.GetProducts)
		productGroup.GET("/:id", productController.GetProductDetail)

		adminGroup := productGroup.Group("/", middleware.Authenticate(), middleware.AuthorizeRole("ADMIN"))
		{
			adminGroup.POST("", productController.CreateProduct)
			adminGroup.PUT("/:id", productController.UpdateProduct)
			adminGroup.POST("/:id/images", productController.AddImages)
			adminGroup.PATCH("/:id/images", productController.UpdateThumbnailProduct)
			adminGroup.DELETE("/images/:id", productController.RemoveImages)
			adminGroup.DELETE("/:id", productController.DeleteProduct)
		}
	}
}
