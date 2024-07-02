package model

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uint       `gorm:"primaryKey"`
	Token     string     `gorm:"unique;not null"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null"`
	ExpiresAt time.Time  `gorm:"not null"`
	Revoked   bool       `gorm:"not null;default:false"`
	CreatedAt *time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (r *RefreshToken) TableName() string {
	return "refresh_token"
}
