package model

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID   uuid.UUID `gorm:"type:uuid;not null"`
	Token    string    `gorm:"type:varchar;not null;index"`
	FamilyID uuid.UUID `gorm:"type:uuid;not null;index"`

	IsUsed    bool      `gorm:"default:false"`
	IsRevoked bool      `gorm:"default:false"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
