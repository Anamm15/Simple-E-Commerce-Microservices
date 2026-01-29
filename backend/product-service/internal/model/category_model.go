package model

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`

	Products []Product `gorm:"many2many:product_categories;" json:"products,omitempty"`
}
