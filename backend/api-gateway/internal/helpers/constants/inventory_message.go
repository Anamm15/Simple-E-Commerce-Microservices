package constants

const (
	// Inventory Success Messages
	MsgStockUpdated   = "Inventory stock updated successfully"
	MsgStockReserved  = "Stock reserved for order"
	MsgRestockSuccess = "Restock operation completed"

	// Inventory Error Messages
	MsgInsufficientStock = "Insufficient stock for the requested quantity"
	MsgStockLockFailed   = "Failed to lock stock for processing"
	MsgWarehouseNotFound = "Warehouse location not found"
	MsgStockUpdateFailed = "Failed to update inventory levels"
)
