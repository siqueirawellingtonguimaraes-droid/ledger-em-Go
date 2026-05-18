package accounts

import "github.com/google/uuid"

func DomainToResponseDTO(account AccountWithBalance) AccountResponseDTO {
	return AccountResponseDTO{
		ID:      account.Account.ID,
		Name:    account.Account.Name,
		Balance: int64(account.Balance.Amount),
	}
}

func CreateDTOToDomain(accountDTO *AccountCreateDTO) Account {
	return Account{
		ID:       uuid.New(),
		Name:     accountDTO.Name,
		Currency: accountDTO.Currency,
	}
}
