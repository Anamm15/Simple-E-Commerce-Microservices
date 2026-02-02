package model

import (
	"time"

	"github.com/google/uuid"
)

type Inventory struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey; default:gen_random_uuid();" json:"id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index:idx_inventory_product_id,unique" json:"product_id" binding:"required"`

	TotalStock     int32 `gorm:"not null;check:total_stock >= 0" json:"total_stock"`
	ReservedStock  int32 `gorm:"not null; default:0" json:"reserved_stock"`
	AvailableStock int32 `gorm:"not null; default:0" json:"available_stock"`

	CreatedAt time.Time `gorm:"not null; default:now()" json:"created_at"`
}
