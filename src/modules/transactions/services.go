package transactions

type TransactionService struct {
	repository TransactionRepository
}

func NewTransactionService(repository TransactionRepository) *TransactionService {
	return &TransactionService{
		repository: repository,
	}
}

func (s *TransactionService) Create(dto *TransactionCreateDTO) error {
	transaction, err := DTOToDomain(dto)
	if err != nil {
		return err
	}

	return s.repository.Save(&transaction)
}
