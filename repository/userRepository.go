package repository

import (
	"github.com/yogawahyudi7/mnc/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByPhoneNumber(phoneNumber string) (*model.User, error)
	GetUserByID(userID uuid.UUID) (*model.User, error)
	UpdateUser(user *model.User) error
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

func (r *userRepository) GetUserByID(userID uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}
