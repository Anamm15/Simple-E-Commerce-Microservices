package model

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	Street   string `gorm:"type:varchar(255);not null" json:"street"`
	City     string `gorm:"type:varchar(100);not null" json:"city"`
	Province string `gorm:"type:varchar(100);not null" json:"province"`
	ZipCode  string `gorm:"type:varchar(10);not null" json:"zip_code"`

	IsPrimary bool `gorm:"default:false" json:"is_primary"`

	UserID uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
