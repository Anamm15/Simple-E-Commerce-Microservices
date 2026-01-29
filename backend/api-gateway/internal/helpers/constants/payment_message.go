package constants

const (
	// Payment Success Messages
	MsgPaymentInitiated = "Payment process started"
	MsgPaymentSuccess   = "Payment received successfully. Thank you!"
	MsgRefundProcessed  = "Refund has been processed successfully"

	// Payment Error Messages
	MsgPaymentFailed      = "Payment transaction failed. Please try again"
	MsgPaymentExpired     = "Payment session has expired"
	MsgInsufficientFunds  = "Your account balance is insufficient"
	MsgUnsupportedMethod  = "The selected payment method is not supported"
	MsgPaymentPending     = "Payment is still being processed by the provider"
	MsgTransactionInvalid = "Transaction details are invalid or corrupted"
)
