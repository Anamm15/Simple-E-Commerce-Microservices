package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`

	Status       string `gorm:"type:varchar(50);not null" json:"status"`
	TotalAmount  int64  `gorm:"not null" json:"total_amount"`
	ShippingCost int64  `gorm:"not null" json:"shipping_cost"`

	ShippingAddressSnapshot string `gorm:"type:varchar(255);not null" json:"shipping_address_snapshot"`

	Items []OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"items"`
	Notes *string     `gorm:"type:text" json:"notes"` // reason for cancellation

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
