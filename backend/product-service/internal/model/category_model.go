package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Name string `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`

	Products []Product `gorm:"many2many:product_categories;" json:"products,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}
