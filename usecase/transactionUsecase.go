package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yogawahyudi7/mnc/config"
	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/repository"
)

type TransactionUsecase interface {
	TopUp(request dto.TopUpRequest, uuid string) (*dto.TopUpResponse, error)
	Transfer(request dto.TransferRequest, uuid string) (*dto.TransferResponse, error)
	Payment(request dto.PaymentRequest) (*dto.PaymentResponse, error)
}

type transactionUsecase struct {
	config    *config.Server
	userRepo  repository.UserRepository
	tokenRepo repository.RefreshTokenRepository
}

func (u *userUsecase) TopUp(request dto.TopUpRequest, uuid string) (*dto.TopUpResponse, error) {
	if request.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	userData, err := u.userRepo.GetUserByID(uuid)
	if err != nil {
		return nil, errors.New("user not found")
	}

	topUpData, err := u.userRepo.TopUp(request.Amount, uuid)
	if err != nil {
		return nil, err
	}

	response := &dto.TopUpResponse{
		TopUpID:       topUpData.Id,
		AmountTopUp:   request.Amount,
		BalanceBefore: userData.Balance,
		BalanceAfter:  topUpData.Balance,
		CreatedDate:   topUpData.CreatedAt,
	}

	return response, nil
}

func (u *userUsecase) Transfer(request dto.TransferRequest, senderId string) (*dto.TransferResponse, error) {
	if request.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	sender, err := u.userRepo.GetUserByPhoneNumber(senderId)
	if err != nil {
		return nil, errors.New("sender not found")
	}

	recipient, err := u.userRepo.GetUserByID(request.TargetUser)
	if err != nil {
		return nil, errors.New("recipient not found")
	}

	if sender.Balance < request.Amount {
		return nil, errors.New("insufficient balance")
	}

	sender.Balance -= request.Amount
	recipient.Balance += request.Amount

	if err := u.userRepo.UpdateUser(sender); err != nil {
		return nil, err
	}

	if err := u.userRepo.UpdateUser(recipient); err != nil {
		return nil, err
	}

	return &dto.TransferResponse{
		TransferID:    uuid.New(),
		Amount:        request.Amount,
		BalanceBefore: sender.Balance + request.Amount,
		BalanceAfter:  sender.Balance,
		CreatedDate:   &u.config.TimeNow,
	}, nil
}

func (u *userUsecase) Payment(request dto.PaymentRequest) (*dto.PaymentResponse, error) {
	return nil, nil
}
