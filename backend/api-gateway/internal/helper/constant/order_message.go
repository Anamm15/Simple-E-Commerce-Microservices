package constant

const (
	// Order Success Messages
	MsgOrderCreated      = "Order placed successfully"
	MsgOrderCancelled    = "Order has been cancelled"
	MsgOrderCompleted    = "Order marked as completed"
	MsgOrderStatusUpdate = "Order status updated successfully"

	// Order Error Messages
	MsgOrderNotFound      = "Order record not found"
	MsgInvalidOrderStatus = "Invalid order status transition"
	MsgOrderAlreadyPaid   = "This order has already been paid"
	MsgCheckoutFailed     = "Failed to process checkout, please try again"
	MsgMinimumOrderValue  = "Order total does not meet the minimum requirement"
)
