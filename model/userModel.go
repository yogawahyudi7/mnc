package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id          uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName   string          `gorm:"column:first_name"`
	LastName    string          `gorm:"column:last_name"`
	PhoneNumber string          `gorm:"column:phone_number;unique"`
	Address     string          `gorm:"column:address"`
	Pin         string          `gorm:"column:pin"`
	Balance     float64         `gorm:"column:balance"`
	CreatedAt   *time.Time      `gorm:"column:created_at"`
	UpdatedAt   *time.Time      `gorm:"column:updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "user"
}
