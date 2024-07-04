package model

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Id        uuid.UUID  `gorm:"type:uuid;not null"`
	Token     string     `gorm:"unique;not null"`
	ExpiresAt time.Time  `gorm:"not null"`
	Revoked   bool       `gorm:"not null;default:false"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (r *RefreshToken) TableName() string {
	return "refresh_token"
}
