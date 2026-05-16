package accounts

func DomainToResponseDTO(account AccountWithBalance) AccountResponseDTO {
	return AccountResponseDTO{
		ID:      account.Account.ID,
		Name:    account.Account.Name,
		Balance: account.Balance,
	}
}

func CreateDTOToDomain(accountDTO AccountCreateDTO) Account {
	return Account{}
}
