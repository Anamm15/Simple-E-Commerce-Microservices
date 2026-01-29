package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	UserID uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`

	Title   string `gorm:"type:varchar(150);not null" json:"title"`
	Message string `gorm:"type:text;not null" json:"message"`

	IsRead bool `gorm:"not null;default:false" json:"is_read"`

	Type string `gorm:"type:varchar(50);index;not null" json:"type"`

	CreatedAt time.Time `json:"created_at"`
}
