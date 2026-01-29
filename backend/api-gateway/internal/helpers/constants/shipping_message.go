package constants

const (
	// Shipping Success Messages
	MsgShippingCalculated = "Shipping cost calculated successfully"
	MsgShipmentCreated    = "Shipment has been prepared"
	MsgTrackingUpdated    = "Tracking information updated"

	// Shipping Error Messages
	MsgAddressInvalid     = "Shipping address is incomplete or invalid"
	MsgCourierUnavailable = "Selected courier service is currently unavailable for your area"
	MsgShipmentNotFound   = "Shipment details could not be found"
	MsgWeightExceeded     = "Total weight exceeds the maximum limit for this service"
)
