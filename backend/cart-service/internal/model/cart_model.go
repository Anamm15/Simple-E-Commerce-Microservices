package model

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`

	UserID    uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	ProductID uuid.UUID `gorm:"type:uuid;index;not null" json:"product_id"`

	Quantity int32 `gorm:"not null;check:quantity > 0" json:"quantity"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
