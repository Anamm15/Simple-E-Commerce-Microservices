package model

import (
	"time"

	"order-service/internal/helper/enum"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Order struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`

	Status       enum.OrderStatus `gorm:"type:varchar(50);not null;default:'PENDING'" json:"status"`
	TotalAmount  int64            `gorm:"not null" json:"total_amount"`
	ShippingCost int64            `gorm:"not null" json:"shipping_cost"`

	ShippingAddressSnapshot datatypes.JSON `gorm:"type:varchar(255);not null" json:"shipping_address_snapshot"`

	Items []OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"items"`
	Notes *string     `gorm:"type:text" json:"notes"` // reason for cancellation

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
