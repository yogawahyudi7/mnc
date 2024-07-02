package repository

import (
	"github.com/yogawahyudi7/mnc/model"

	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	CreateRefreshToken(token *model.RefreshToken) error
	GetToken(token string) (*model.RefreshToken, error)
	DeleteToken(token string) error
	RevokeToken(token string) error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &tokenRepository{db}
}

func (r *tokenRepository) CreateRefreshToken(token *model.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *tokenRepository) GetToken(token string) (*model.RefreshToken, error) {
	var refreshToken model.RefreshToken
	if err := r.db.Where("token = ?", token).First(&refreshToken).Error; err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *tokenRepository) DeleteToken(token string) error {
	return r.db.Delete(&model.RefreshToken{}, "token = ?", token).Error
}

func (r *tokenRepository) RevokeToken(token string) error {
	return r.db.Model(&model.RefreshToken{}).Where("token = ?", token).Update("revoked", true).Error
}
