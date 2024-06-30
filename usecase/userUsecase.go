package usecase

import (
	"errors"

	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/model"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/pkg/validator"
	"github.com/yogawahyudi7/mnc/repository"

	"github.com/google/uuid"
)

type UserUsecase interface {
	RegisterUser(request dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(request dto.LoginRequest) (*dto.LoginResponse, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) RegisterUser(request dto.RegisterRequest) (*dto.RegisterResponse, error) {

	// Validasi PIN
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

func (u *userUsecase) Login(request dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := u.userRepo.GetUserByPhoneNumber(request.PhoneNumber)
	if err != nil || user.Pin != request.Pin {
		return nil, errors.New("phone number and PIN doesn't match")
	}

	// Generate JWT token (simplified for brevity)
	token := "generated-jwt-token"
	refreshToken := "generated-refresh-token"

	response := &dto.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	return response, nil
}
