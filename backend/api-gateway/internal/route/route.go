package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// Controllers
	controller "api-gateway/internal/controller"

	// PB Clients
	authpb "api-gateway/internal/pb/auth"
	cartpb "api-gateway/internal/pb/cart"
	inventorypb "api-gateway/internal/pb/inventory"
	notificationpb "api-gateway/internal/pb/notification"
	orderpb "api-gateway/internal/pb/order"
	paymentpb "api-gateway/internal/pb/payment"
	productpb "api-gateway/internal/pb/product"
	shippingpb "api-gateway/internal/pb/shipping"
	userpb "api-gateway/internal/pb/user"
)

func SetupRouter(
	authClient authpb.AuthServiceClient,
	userClient userpb.UserServiceClient,
	productClient productpb.ProductServiceClient,
	inventoryClient inventorypb.InventoryServiceClient,
	cartClient cartpb.CartServiceClient,
	orderClient orderpb.OrderServiceClient,
	paymentClient paymentpb.PaymentServiceClient,
	shippingClient shippingpb.ShippingServiceClient,
	notificationClient notificationpb.NotificationServiceClient,
) *gin.Engine {
	r := gin.Default()

	// Global Middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	api := r.Group("/api/v1")
	// Inisialisasi Controller
	authCtrl := controller.NewAuthController(authClient)
	userCtrl := controller.NewUserController(userClient)
	productCtrl := controller.NewProductController(productClient)
	inventoryCtrl := controller.NewInventoryController(inventoryClient)
	cartCtrl := controller.NewCartController(cartClient)
	orderCtrl := controller.NewOrderController(orderClient)
	paymentCtrl := controller.NewPaymentController(paymentClient)
	notificationCtrl := controller.NewNotificationController(notificationClient)

	// Registrasi Route per Domain
	RegisterAuthRoutes(api, authCtrl)
	RegisterUserRoutes(api, userCtrl)
	RegisterProductRoutes(api, productCtrl)
	RegisterInventoryRoutes(api, inventoryCtrl)
	RegisterCartRoutes(api, cartCtrl)
	RegisterOrderRoutes(api, orderCtrl)
	RegisterPaymentRoutes(api, paymentCtrl)
	RegisterNotificationRoutes(api, notificationCtrl)

	return r
}
