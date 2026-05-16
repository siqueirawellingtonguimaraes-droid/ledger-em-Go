package accounts

import (
	"errors"
	vo "ledger/src/VO"
	"strings"

	"github.com/google/uuid"
)

type Account struct {
	ID       uuid.UUID
	Name     string
	Currency string
}

type AccountWithBalance struct {
	Account Account
	Balance vo.Money
}

func NewAccount(name, currency string) (Account, error) {
	if len(strings.TrimSpace(name)) <= 0 {
		return Account{}, errors.New("O nome deve ser preenchido")
	}

	return Account{
		ID:       uuid.New(),
		Name:     name,
		Currency: currency,
	}, nil
}
