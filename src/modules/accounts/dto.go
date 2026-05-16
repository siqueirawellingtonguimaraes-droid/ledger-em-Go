package accounts

import (
	vo "ledger/src/VO"

	"github.com/google/uuid"
)

type AccountCreateDTO struct {
	Name     string `json:"name" validate:"required"`
	Currency string `json:"currency"`
}

type AccountResponseDTO struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Balance vo.Money  `json:"balance"`
}
