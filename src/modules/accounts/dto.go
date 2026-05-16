package accounts

import (
	vo "ledger/src/VO"

	"github.com/google/uuid"
)

type AccountCreateDTO struct {
	Name     string      `json:"name" validate:"required"`
	Currency vo.Currency `json:"currency"`
}

type AccountResponseDTO struct {
	ID       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	Currency vo.Currency `json:"currency"`
	Balance  vo.Money    `json:"balance"`
}
