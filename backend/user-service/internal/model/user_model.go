package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	FullName    string `gorm:"type:varchar(100);not null" json:"full_name"`
	PhoneNumber string `gorm:"type:varchar(20);uniqueIndex" json:"phone_number"`

	Addresses []Address `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"addresses"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
