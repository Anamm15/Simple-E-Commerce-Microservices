package model

import (
	"github.com/google/uuid"
)

type CartItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	UserID    uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	ProductID uuid.UUID `gorm:"type:uuid;index;not null" json:"product_id"`

	Quantity int32 `gorm:"not null;check:quantity > 0" json:"quantity"`
}
