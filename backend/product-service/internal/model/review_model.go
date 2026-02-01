package model

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	ProductID uuid.UUID `gorm:"type:uuid;index;not null" json:"product_id"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`

	UserName string `gorm:"type:varchar(100);not null" json:"user_name"`
	Rating   int32  `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Comment  string `gorm:"type:text" json:"comment"`

	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE" json:"-"`

	CreatedAt time.Time `json:"created_at"`
}
