package accounts

import (
	vo "ledger/src/VO"
)

type AccountService struct {
	repository AccountRepository
}

func NewAccountService(repository AccountRepository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}

func (s *AccountService) Create(dto *AccountCreateDTO) error {
	account := CreateDTOToDomain(dto)
	return s.repository.Save(&account)
}

func (s *AccountService) GetByID(id string) (AccountResponseDTO, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return AccountResponseDTO{}, err
	}

	accountWithBalance := AccountWithBalance{
		Account: *account,
		Balance: calculateBalance(),
	}

	return DomainToResponseDTO(accountWithBalance), nil
}

func calculateBalance() vo.Money {
	return vo.Money{
		Amount:   150000,
		Currency: vo.BRL,
	}
}
