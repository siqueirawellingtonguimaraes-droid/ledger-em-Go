package accounts

import (
	vo "ledger/src/VO"
	"ledger/src/modules/transactions"
)

type AccountService struct {
	repository      AccountRepository
	transactionRepo transactions.TransactionRepository
}

func NewAccountService(repository AccountRepository, transactionRepo transactions.TransactionRepository) *AccountService {
	return &AccountService{
		repository:      repository,
		transactionRepo: transactionRepo,
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

	balance, err := s.calculateBalance(id, account.Currency)
	if err != nil {
		return AccountResponseDTO{}, err
	}

	accountWithBalance := AccountWithBalance{
		Account: *account,
		Balance: balance,
	}

	return DomainToResponseDTO(accountWithBalance), nil
}

func (s *AccountService) calculateBalance(accountId string, currency vo.Currency) (vo.Money, error) {
	entries, err := s.transactionRepo.GetEntriesByAccountID(accountId)
	if err != nil {
		return vo.Money{}, err
	}

	var total int64

	for _, entry := range entries {
		if entry.Amount.Currency != currency {
			return vo.Money{}, nil
		}

		amount := int64(entry.Amount.Amount)
		if entry.Direction == vo.Debit {
			total -= amount
		} else {
			total += amount
		}
	}

	return vo.Money{
		Amount:   uint64(total),
		Currency: currency,
	}, nil
}
