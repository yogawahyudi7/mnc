package repository

import (
	"github.com/yogawahyudi7/mnc/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByPhoneNumber(phoneNumber string) (*model.User, error)
	GetUserByID(userId string) (*model.User, error)
	UpdateUser(user *model.User) error
	TopUp(amount float64, uuid string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByPhoneNumber(phoneNumber string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(userId string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *model.User) error {
	return r.db.Where("id = ?", user.Id).Updates(user).Error
}

func (r *userRepository) TopUp(amount float64, uuid string) error {
	var user model.User
	if err := r.db.Model(&user).Where("id = ?", uuid).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		return err
	}
	return nil
}
