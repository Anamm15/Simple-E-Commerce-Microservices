package enum

type InventoryReservationStatus string

const (
	InventoryReservationStatusReserved InventoryReservationStatus = "RESERVED"
	InventoryReservationStatusReleased InventoryReservationStatus = "RELEASED"
	InventoryReservationStatusCommited InventoryReservationStatus = "COMMITED"
)
