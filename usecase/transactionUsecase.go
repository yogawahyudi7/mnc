package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yogawahyudi7/mnc/config"
	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/model"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/repository"
)

type TransactionUsecase interface {
	TopUp(request dto.TopUpRequest, id string) (*dto.TopUpResponse, error)
	Transfer(request dto.TransferRequest, id string) (*dto.TransferResponse, error)
	Payment(request dto.PaymentRequest, id string) (*dto.PaymentResponse, error)
	ListTransactions(id string) ([]dto.TransactionHistoryResponse, error)
}

type transactionUsecase struct {
	config          *config.Server
	userRepo        repository.UserRepository
	transactionRepo repository.TransactionRepository
}

func NewTransactionUsecase(config *config.Server, userRepo repository.UserRepository, transactionRepo repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		config:          config,
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}

func (u *transactionUsecase) TopUp(request dto.TopUpRequest, id string) (*dto.TopUpResponse, error) {
	if request.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	userData, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = u.transactionRepo.TopUp(request.Amount, id)
	if err != nil {
		return nil, err
	}

	dataTransaction := &model.Transaction{
		Id:              uuid.New(),
		UserId:          userData.Id,
		Amount:          request.Amount,
		TransactionType: constant.TopUpType,
		BalanceBefore:   userData.Balance,
		BalanceAfter:    userData.Balance + request.Amount,
		Remarks:         constant.TopUpRemarks,
	}

	err = u.transactionRepo.CreateTransaction(dataTransaction)
	if err != nil {
		return nil, err
	}

	response := &dto.TopUpResponse{
		TopUpID:       dataTransaction.Id,
		AmountTopUp:   request.Amount,
		BalanceBefore: userData.Balance,
		BalanceAfter:  dataTransaction.BalanceAfter,
		CreatedDate:   dataTransaction.CreatedAt.Format(constant.TimeFormatYMDHMS),
	}

	return response, nil
}

func (u *transactionUsecase) Transfer(request dto.TransferRequest, senderId string) (*dto.TransferResponse, error) {
	if request.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	sender, err := u.userRepo.GetUserByID(senderId)
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

	err = u.transactionRepo.Transfer(request.Amount, sender.Id.String(), recipient.Id.String())
	if err != nil {
		return nil, err
	}

	dataTransaction := &model.Transaction{
		Id:              uuid.New(),
		UserId:          sender.Id,
		Amount:          request.Amount,
		TransactionType: constant.TransferType,
		BalanceBefore:   sender.Balance + request.Amount,
		BalanceAfter:    sender.Balance,
		Remarks:         request.Remarks,
	}
	err = u.transactionRepo.CreateTransaction(dataTransaction)
	if err != nil {
		return nil, err
	}

	return &dto.TransferResponse{
		TransferID:    uuid.New(),
		Amount:        request.Amount,
		BalanceBefore: sender.Balance + request.Amount,
		BalanceAfter:  sender.Balance,
		Remarks:       request.Remarks,
		CreatedDate:   time.Now().Format(constant.TimeFormatYMDHMS),
	}, nil
}

func (u *transactionUsecase) Payment(request dto.PaymentRequest, id string) (*dto.PaymentResponse, error) {
	if request.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	userData, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if userData.Balance < request.Amount {
		return nil, errors.New("insufficient balance")
	}

	err = u.transactionRepo.Payment(request.Amount, userData.Id.String())
	if err != nil {
		return nil, err
	}

	dataTransaction := &model.Transaction{
		Id:              uuid.New(),
		UserId:          userData.Id,
		Amount:          request.Amount,
		TransactionType: constant.PaymentType,
		BalanceBefore:   userData.Balance,
		BalanceAfter:    userData.Balance - request.Amount,
		Remarks:         request.Remarks,
	}

	err = u.transactionRepo.CreateTransaction(dataTransaction)
	if err != nil {
		return nil, err
	}

	response := &dto.PaymentResponse{
		PaymentID:     dataTransaction.Id,
		AmountPayment: request.Amount,
		BalanceBefore: userData.Balance,
		BalanceAfter:  dataTransaction.BalanceAfter,
		Remarks:       request.Remarks,
		CreatedDate:   dataTransaction.CreatedAt.Format(constant.TimeFormatYMDHMS),
	}

	return response, nil
}

func (u *transactionUsecase) ListTransactions(id string) ([]dto.TransactionHistoryResponse, error) {
	transactions, err := u.transactionRepo.GetTransactionByUserID(id)
	if err != nil {
		return nil, err
	}

	var response []dto.TransactionHistoryResponse
	for _, transaction := range transactions {
		response = append(response, dto.TransactionHistoryResponse{
			TransactionID:   transaction.Id,
			TransactionType: transaction.TransactionType,
			UserId:          transaction.UserId,
			Amount:          transaction.Amount,
			Remarks:         transaction.Remarks,
			BalanceBefore:   transaction.BalanceBefore,
			BalanceAfter:    transaction.BalanceAfter,
			CreatedDate:     transaction.CreatedAt.Format(constant.TimeFormatYMDHMS),
		})
	}

	return response, nil
}
