package accounts

import (
	vo "ledger/src/VO"

	"github.com/google/uuid"
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
	account := Account{
		ID:       uuid.New(),
		Name:     dto.Name,
		Currency: dto.Currency,
	}
	return s.repository.Save(&account)
}

func (s *AccountService) GetByID(id string) (*AccountResponseDTO, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	accountWithBalance := AccountWithBalance{
		Account: *account,
		Balance: vo.Money{Amount: 15000, Currency: account.Currency},
	}

	return &AccountResponseDTO{
		ID:       accountWithBalance.Account.ID,
		Name:     accountWithBalance.Account.Name,
		Currency: accountWithBalance.Account.Currency,
		Balance:  accountWithBalance.Balance,
	}, nil
}
