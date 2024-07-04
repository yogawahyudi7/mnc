package dto

import (
	"github.com/google/uuid"
)

type RegisterRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Pin         string `json:"pin"`
}

type RegisterResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedDate string    `json:"created_date"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"pin"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TopUpRequest struct {
	Amount float64 `json:"amount"`
}

type TopUpResponse struct {
	TopUpID       uuid.UUID `json:"top_up_id"`
	AmountTopUp   float64   `json:"amount_top_up"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedDate   string    `json:"created_date"`
}

type TransferRequest struct {
	TargetUser string  `json:"target_user"`
	Amount     float64 `json:"amount"`
	Remarks    string  `json:"remarks"`
}

type TransferResponse struct {
	TransferID    uuid.UUID `json:"transfer_id"`
	Amount        float64   `json:"amount"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedDate   string    `json:"created_date"`
}

type PaymentRequest struct {
	Amount  float64 `json:"amount"`
	Remarks string  `json:"remarks"`
}

type PaymentResponse struct {
	PaymentID     uuid.UUID `json:"payment_id"`
	AmountPayment float64   `json:"amount_payment"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	Remarks       string    `json:"remarks"`
	CreatedDate   string    `json:"created_date"`
}

type TransactionHistoryResponse struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	Transaction   string    `json:"transaction"`
	Amount        float64   `json:"amount"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedDate   string    `json:"created_date"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

type UpdateUserResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	UpdatedDate string    `json:"updated_date"`
}

type GetUserResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedDate string    `json:"created_date"`
}
