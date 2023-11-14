package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"not null"`
	CreatedAt time.Time
	ExpiresAt time.Time
}
