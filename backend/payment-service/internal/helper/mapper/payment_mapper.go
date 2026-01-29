package mapper

import "payment-service/internal/helper/enum"

func MapMidtransStatus(
	transactionStatus string,
	fraudStatus string,
) (enum.PaymentStatus, enum.OrderStatus) {
	switch transactionStatus {

	case "pending":
		return enum.PaymentPending, enum.OrderPending

	case "capture":
		if fraudStatus == "challenge" {
			return enum.PaymentAuthorized, enum.OrderPending
		}
		if fraudStatus == "accept" {
			return enum.PaymentPaid, enum.OrderPaid
		}
		return enum.PaymentFailed, enum.OrderCancelled

	case "settlement":
		return enum.PaymentPaid, enum.OrderPaid

	case "deny":
		return enum.PaymentFailed, enum.OrderCancelled

	case "cancel":
		return enum.PaymentCancelled, enum.OrderCancelled

	case "expire":
		return enum.PaymentExpired, enum.OrderExpired

	case "refund":
		return enum.PaymentRefunded, enum.OrderRefunded

	case "partial_refund":
		return enum.PaymentPartiallyRefunded, enum.OrderRefunded
	}

	return enum.PaymentPending, enum.OrderPending
}
