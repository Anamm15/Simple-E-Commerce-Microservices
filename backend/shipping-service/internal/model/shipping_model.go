package model

import (
	"time"

	"github.com/google/uuid"
)

type Shipment struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID uuid.UUID `gorm:"type:uuid;index;not null" json:"order_id"`

	CourierCode string `gorm:"type:varchar(50);not null" json:"courier_code"`
	ServiceCode string `gorm:"type:varchar(50);not null" json:"service_code"`

	TrackingNumber    string    `gorm:"type:varchar(100);uniqueIndex" json:"tracking_number"`
	Status            string    `gorm:"type:varchar(30);index;not null" json:"status"`
	ShippingCost      int64     `gorm:"not null" json:"shipping_cost"`
	EstimatedDelivery time.Time `gorm:"not null" json:"estimated_delivery"`

	CreatedAt time.Time `json:"created_at"`
}
