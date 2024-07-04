package usecase

import (
	"github.com/yogawahyudi7/mnc/config"
	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/repository"
)

type UserUsecase interface {
	UpdateUser(request dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
	GetUser(uuid string) (*dto.GetUserResponse, error)
}

type userUsecase struct {
	config    *config.Server
	userRepo  repository.UserRepository
	tokenRepo repository.RefreshTokenRepository
}

func NewUserUsecase(config *config.Server, userRepo repository.UserRepository, tokenRepo repository.RefreshTokenRepository) UserUsecase {
	return &userUsecase{config, userRepo, tokenRepo}
}

func (u *userUsecase) UpdateUser(request dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {

	return nil, nil
}

func (u *userUsecase) GetUser(uuid string) (*dto.GetUserResponse, error) {

	user, err := u.userRepo.GetUserByID(uuid)
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
