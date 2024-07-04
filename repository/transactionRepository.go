package repository

import (
	"github.com/yogawahyudi7/mnc/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *model.Transaction) error
	GetTransactionByUserID(id string) ([]model.Transaction, error)
	TopUp(amount float64, uuid string) error
	Transfer(amount float64, senderId, receiverId string) error
	Payment(amount float64, id string) error
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

func (r *transactionRepository) GetTransactionByUserID(id string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := r.db.Where("user_id = ?", id).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) TopUp(amount float64, uuid string) error {
	var user model.User
	if err := r.db.Model(&user).Where("id = ?", uuid).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) Transfer(amount float64, senderId, receiverId string) error {
	tx := r.db.Begin()
	if err := tx.Model(&model.User{}).Where("id = ?", senderId).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.User{}).Where("id = ?", receiverId).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *transactionRepository) Payment(amount float64, id string) error {
	if err := r.db.Model(&model.User{}).Where("id = ?", id).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		return err
	}
	return nil
}
