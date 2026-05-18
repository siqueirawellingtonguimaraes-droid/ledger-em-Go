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
	Currency vo.Currency
}

type AccountWithBalance struct {
	Account Account
	Balance vo.Money
}

func NewAccount(name string, currency vo.Currency) (Account, error) {
	if len(strings.TrimSpace(name)) <= 0 {
		return Account{}, errors.New("O nome deve ser preenchido")
	}

	return Account{
		ID:       uuid.New(),
		Name:     name,
		Currency: currency,
	}, nil
}

func NewAccountWithBalance(account Account) (AccountWithBalance, error) {

	balance, err := vo.NewMoney(0, account.Currency)

	if err != nil {
		return AccountWithBalance{}, err
	}

	return AccountWithBalance{
		Account: account,
		Balance: balance,
	}, nil
}
