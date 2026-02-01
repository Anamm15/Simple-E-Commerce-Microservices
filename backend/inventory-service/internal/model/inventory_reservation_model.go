package model

import (
	"time"

	"inventory-service/internal/helper/enum"

	"github.com/google/uuid"
)

type InventoryReservation struct {
	ID          uint64                          `gorm:"primaryKey;autoIncrement" json:"id"`
	InventoryID uuid.UUID                       `gorm:"type:uuid;not null;index:idx_inventory_reservation_inventory_id,unique" json:"inventory_id" binding:"required"`
	ProductID   uuid.UUID                       `gorm:"type:uuid;not null;index:idx_inventory_reservation_product_id,unique" json:"product_id" binding:"required"`
	OrderID     uuid.UUID                       `gorm:"type:uuid;not null;index:idx_inventory_reservation_order_id,unique" json:"order_id" binding:"required"`
	Quantity    int32                           `gorm:"not null;check:quantity > 0" json:"quantity" binding:"required"`
	Status      enum.InventoryReservationStatus `gorm:"not null;default:'RESERVED'" json:"status"`
	ExpiryAt    time.Time                       `gorm:"not null" json:"expiry_at"`
	CreatedAt   time.Time                       `gorm:"not null; default:now()" json:"created_at"`
}
