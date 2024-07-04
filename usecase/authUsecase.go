package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yogawahyudi7/mnc/config"
	"github.com/yogawahyudi7/mnc/dto"
	model "github.com/yogawahyudi7/mnc/model"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/pkg/jwt"
	"github.com/yogawahyudi7/mnc/pkg/validator"
	"github.com/yogawahyudi7/mnc/repository"
)

type AuthUsecase interface {
	RegisterUser(request dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(request dto.LoginRequest) (*dto.LoginResponse, error)
	RefreshToken(token string) (*dto.LoginResponse, error)
}

type authUsecase struct {
	config    *config.Server
	userRepo  repository.UserRepository
	tokenRepo repository.RefreshTokenRepository
}

func NewAuthUsecase(config *config.Server, userRepo repository.UserRepository, tokenRepo repository.RefreshTokenRepository) AuthUsecase {
	return &authUsecase{config, userRepo, tokenRepo}
}

func (u *authUsecase) RegisterUser(request dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if !validator.IsValidPin(request.Pin) {
		return nil, errors.New("PIN must be a 6-digit number")
	}

	existingUser, _ := u.userRepo.GetUserByPhoneNumber(request.PhoneNumber)
	if existingUser != nil {
		return nil, errors.New("phone number already registered")
	}

	user := &model.User{
		Id:          uuid.New(),
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		Pin:         request.Pin,
	}

	if err := u.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		UserID:      user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		CreatedDate: user.CreatedAt.Format(constant.TimeFormatYMDHMS),
	}

	return response, nil
}

func (u *authUsecase) Login(request dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.GetUserByPhoneNumber(request.PhoneNumber)
	if err != nil || user.Pin != request.Pin {
		return nil, errors.New("phone number and PIN doesn't match")
	}

	token, err := jwt.GenerateToken(u.config, user.Id.String(), constant.TokenTypeAccess, u.config.TokenDuration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateToken(u.config, user.Id.String(), constant.TokenTypeRefresh, u.config.RefreshTokenDuration)
	if err != nil {
		return nil, err
	}

	timeDuration, _ := time.ParseDuration(u.config.RefreshTokenDuration)

	refreshTokenModel := &model.RefreshToken{
		Id:        user.Id,
		Token:     refreshToken,
		Revoked:   false,
		ExpiresAt: time.Now().Add(timeDuration),
	}
	if err := u.tokenRepo.CreateRefreshToken(refreshTokenModel); err != nil {
		return nil, err
	}

	response := &dto.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (u *authUsecase) RefreshToken(token string) (*dto.LoginResponse, error) {
	refreshToken, err := u.tokenRepo.GetToken(token)
	if err != nil || refreshToken.ExpiresAt.Before(time.Now()) || refreshToken.Revoked {
		return nil, errors.New(constant.InvalidExpiredToken)
	}

	newAccessToken, err := jwt.GenerateToken(u.config, refreshToken.Id.String(), constant.TokenTypeAccess, u.config.TokenDuration)
	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponse{
		AccessToken:  newAccessToken,
		RefreshToken: token, // Refresh token remains the same
	}

	return response, nil
}
