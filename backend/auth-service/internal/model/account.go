package model

import (
	"time"

	"auth-service/internal/helper/enum"

	"github.com/google/uuid"
)

type Account struct {
	ID         uuid.UUID        `gorm:"primaryKey;autoIncrement;default:gen_random_uuid()"`
	UserID     uuid.UUID        `gorm:"uniqueIndex;not null"`
	Email      string           `gorm:"uniqueIndex;not null"`
	Username   string           `gorm:"uniqueIndex;not null"`
	Password   string           `gorm:"not null"`
	Role       enum.AccountRole `gorm:"not null default:'user'"`
	IsVerified bool             `gorm:"not null default:false"`
	CreatedAt  time.Time        `gorm:"autoCreateTime"`
}
