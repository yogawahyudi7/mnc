package repository

import (
	"github.com/yogawahyudi7/mnc/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *model.Transaction) error
	GetTransactionByUserID(userId string) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(transaction *model.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetTransactionByUserID(userId string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := r.db.Where("user_id = ?", userId).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
