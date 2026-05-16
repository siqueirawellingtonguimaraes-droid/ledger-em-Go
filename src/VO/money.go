package vo

import (
	"errors"
)

type Money struct {
	Amount   uint64
	Currency Currency
}

func NewMoney(amount uint64, currency Currency) (Money, error) {
	if currency == "" {
		return Money{}, errors.New("currency is required")
	}

	return Money{
		Amount:   amount,
		Currency: currency,
	}, nil
}

func (m Money) GetAmount() uint64 {
	return m.Amount
}

func (m Money) GetCurrency() Currency {
	return m.Currency
}

func (m Money) checkCurrency(other Money) error {
	if m.Currency != other.Currency {
		return errors.New("currencies do not match")
	}
	return nil
}

func (m Money) Add(other Money) (Money, error) {
	if err := m.checkCurrency(other); err != nil {
		return Money{}, err
	}

	return Money{
		Amount:   m.Amount + other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m Money) Subtract(other Money) (Money, error) {
	if err := m.checkCurrency(other); err != nil {
		return Money{}, err
	}

	return Money{
		Amount:   m.Amount - other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m Money) Equal(other Money) bool {
	return m.Amount == other.Amount && m.Currency == other.Currency
}
