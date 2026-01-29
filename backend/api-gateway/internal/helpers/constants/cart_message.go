package constants

const (
	// Cart Success Messages
	MsgItemAddedToCart = "Item successfully added to cart"
	MsgCartUpdated     = "Cart quantity updated successfully"
	MsgItemRemoved     = "Item removed from cart"
	MsgCartCleared     = "Cart has been emptied"

	// Cart Error Messages
	MsgCartNotFound       = "Cart session not found"
	MsgItemNotFoundInCart = "The item is no longer in your cart"
	MsgExceedMaxQuantity  = "You have reached the maximum limit for this item"
	MsgCartSyncFailed     = "Failed to sync cart with current inventory"
)
