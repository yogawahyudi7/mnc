package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	Id              uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserId          uuid.UUID       `gorm:"column:user_id"`
	Amount          float64         `gorm:"column:amount"`
	TransactionType string          `gorm:"column:transaction_type"`
	BalanceBefore   float64         `gorm:"column:balance_before"`
	BalanceAfter    float64         `gorm:"column:balance_after"`
	CreatedAt       *time.Time      `gorm:"column:created_at"`
	UpdatedAt       *time.Time      `gorm:"column:updated_at"`
	DeletedAt       *gorm.DeletedAt `gorm:"column:deleted_at"`
}
