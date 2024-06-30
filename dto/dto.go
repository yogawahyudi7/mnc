package dto

import (
	"time"

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
	TopUpID       uuid.UUID  `json:"top_up_id"`
	AmountTopUp   float64    `json:"amount_top_up"`
	BalanceBefore float64    `json:"balance_before"`
	BalanceAfter  float64    `json:"balance_after"`
	CreatedDate   *time.Time `json:"created_date"`
}
