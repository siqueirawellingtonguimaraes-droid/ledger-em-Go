package transactions

import (
	"errors"
	vo "ledger/src/VO"

	"github.com/google/uuid"
)

type TransactionEntry struct {
	ID        uuid.UUID
	AccountID string
	Amount    vo.Money
	Direction vo.Direction
}

type Transaction struct {
	ID          uuid.UUID
	Description string
	Entries     []TransactionEntry
}

func NewTransaction(description string, entries []TransactionEntry) (Transaction, error) {
	if len(entries) < 2 {
		return Transaction{}, errors.New("a transaction must have at least two entries")
	}

	if !isBalanced(entries) {
		return Transaction{}, errors.New("transaction is not balanced")
	}

	return Transaction{
		ID:          uuid.New(),
		Description: description,
		Entries:     entries,
	}, nil
}

func isBalanced(entries []TransactionEntry) bool {
	var total int64

	for _, entry := range entries {
		if entry.Direction == vo.Debit {
			total += int64(entry.Amount.Amount)
		} else {
			total -= int64(entry.Amount.Amount)
		}
	}

	if total != 0 {
		return false
	}
	return true
}
