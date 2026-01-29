package enum

type OrderStatus string

const (
	// Order just created, waiting for payment
	OrderPending OrderStatus = "pending"

	// Payment received / authorized
	OrderPaid OrderStatus = "paid"

	// Seller/admin confirmed the order
	OrderConfirmed OrderStatus = "confirmed"

	// Order is being prepared (packing, processing)
	OrderProcessing OrderStatus = "processing"

	// Order handed over to courier
	OrderShipped OrderStatus = "shipped"

	// Order delivered to customer
	OrderDelivered OrderStatus = "delivered"

	// Order successfully completed (post-delivery success)
	OrderCompleted OrderStatus = "completed"

	// Order cancelled before shipment
	OrderCancelled OrderStatus = "cancelled"

	// Order expired (payment timeout)
	OrderExpired OrderStatus = "expired"

	// Order returned by customer
	OrderReturned OrderStatus = "returned"

	// Refund process completed
	OrderRefunded OrderStatus = "refunded"
)
