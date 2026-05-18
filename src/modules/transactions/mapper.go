package transactions

import (
	"fmt"
	vo "ledger/src/VO"

	"github.com/google/uuid"
)

func DTOToDomain(dto *TransactionCreateDTO) (Transaction, error) {
	entries := make([]TransactionEntry, 0, len(dto.Entries))

	for _, entryDTO := range dto.Entries {
		money, err := vo.NewMoney(
			uint64(entryDTO.Amount),
			vo.Currency(entryDTO.Currency),
		)
		if err != nil {
			return Transaction{}, err
		}

		var direction vo.Direction

		switch entryDTO.Direction {
		case "debit":
			direction = vo.Debit

		case "credit":
			direction = vo.Credit

		default:
			return Transaction{}, fmt.Errorf(
				"invalid direction: %s",
				entryDTO.Direction,
			)
		}

		entry := TransactionEntry{
			ID:        uuid.New(),
			AccountID: entryDTO.AccountID,
			Amount:    money,
			Direction: direction,
		}

		entries = append(entries, entry)
	}

	return NewTransaction(
		dto.Description,
		entries,
	)
}

func DomainToDTO(transaction Transaction) TransactionResponseDTO {
	entries := make([]TransactionEntryResponseDTO, 0, len(transaction.Entries))

	for _, entry := range transaction.Entries {
		var direction vo.Direction

		switch entry.Direction {
		case vo.Debit:
			direction = vo.Debit

		case vo.Credit:
			direction = vo.Credit
		}

		entries = append(
			entries,
			TransactionEntryResponseDTO{
				ID:        entry.ID.String(),
				AccountID: entry.AccountID,
				Amount:    int64(entry.Amount.Amount),
				Currency:  string(entry.Amount.Currency),
				Direction: string(direction),
			},
		)
	}

	return TransactionResponseDTO{
		ID:          transaction.ID.String(),
		Description: transaction.Description,
		Entries:     entries,
	}
}
