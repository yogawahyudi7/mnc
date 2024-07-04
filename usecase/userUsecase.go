package usecase

import (
	"github.com/google/uuid"
	"github.com/yogawahyudi7/mnc/config"
	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/model"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/repository"
)

type UserUsecase interface {
	UpdateUser(request dto.UpdateUserRequest, userId string) (*dto.UpdateUserResponse, error)
	GetUser(id string) (*dto.GetUserResponse, error)
}

type userUsecase struct {
	config    *config.Server
	userRepo  repository.UserRepository
	tokenRepo repository.RefreshTokenRepository
}

func NewUserUsecase(config *config.Server, userRepo repository.UserRepository, tokenRepo repository.RefreshTokenRepository) UserUsecase {
	return &userUsecase{config, userRepo, tokenRepo}
}

func (u *userUsecase) UpdateUser(request dto.UpdateUserRequest, userId string) (*dto.UpdateUserResponse, error) {

	uuid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	updateUser := &model.User{
		Id:        uuid,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.Address,
	}

	err = u.userRepo.UpdateUser(updateUser)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	userResponse := dto.UpdateUserResponse{
		UserID:      user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}

	return &userResponse, nil
}

func (u *userUsecase) GetUser(id string) (*dto.GetUserResponse, error) {

	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	userResponse := dto.GetUserResponse{
		UserID:      user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		CreatedDate: user.CreatedAt.Format(constant.TimeFormatYMDHMS),
	}

	return &userResponse, nil
}
