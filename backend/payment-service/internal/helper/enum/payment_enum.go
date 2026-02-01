package enum

type PaymentStatus string

const (
	// Payment created, waiting for customer action
	PaymentPending PaymentStatus = "pending"

	// Payment authorized (card authorized but not captured)
	PaymentAuthorized PaymentStatus = "authorized"

	// Payment successfully captured / settled
	PaymentPaid PaymentStatus = "paid"

	// Payment failed (insufficient funds, rejected, etc)
	PaymentFailed PaymentStatus = "failed"

	// Payment expired (timeout, no action)
	PaymentExpired PaymentStatus = "expired"

	// Payment cancelled by user or system
	PaymentCancelled PaymentStatus = "cancelled"

	// Refund requested but not finished
	PaymentRefundPending PaymentStatus = "refund_pending"

	// Refund completed
	PaymentRefunded PaymentStatus = "refunded"

	// Partial refund completed
	PaymentPartiallyRefunded PaymentStatus = "partially_refunded"
)
