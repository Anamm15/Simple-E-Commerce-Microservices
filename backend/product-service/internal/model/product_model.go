package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Name        string `gorm:"type:varchar(150);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	Price     int64  `gorm:"not null" json:"price"`
	WeightG   int32  `gorm:"not null" json:"weight_g"`
	Thumbnail string `gorm:"type:text" json:"thumbnail"`

	Images     []Image    `gorm:"foreignKey:ProductID" json:"images"`
	Reviews    []Review   `gorm:"foreignKey:ProductID" json:"reviews"`
	Categories []Category `gorm:"many2many:product_categories;" json:"categories"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
