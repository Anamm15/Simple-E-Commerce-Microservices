package model

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	ProductID uuid.UUID `gorm:"type:uuid;index;not null" json:"product_id"`
	URL       string    `gorm:"type:text;not null" json:"url"`
	PublicID  string    `gorm:"type:text;not null" json:"public_id"`

	Product   Product   `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
