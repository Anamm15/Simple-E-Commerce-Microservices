package model

import "github.com/google/uuid"

type Inventory struct {
	ProductID uuid.UUID `gorm:"type:uuid;primaryKey" json:"product_id" binding:"required"`

	Stock int32 `gorm:"not null;check:stock >= 0" json:"stock"`
}
