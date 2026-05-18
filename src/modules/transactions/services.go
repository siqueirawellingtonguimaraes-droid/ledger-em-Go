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

func (s *TransactionService) GetByID(id string) (TransactionResponseDTO, error) {
	transaction, err := s.repository.FindByID(id)
	if err != nil {
		return TransactionResponseDTO{}, err
	}

	return DomainToDTO(transaction), nil
}

func (s *TransactionService) GetAll() ([]TransactionResponseDTO, error) {
	transactions, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]TransactionResponseDTO, 0, len(transactions))
	for _, transaction := range transactions {
		response = append(response, DomainToDTO(transaction))
	}

	return response, nil
}

func (s *TransactionService) GetEntriesByAccountID(accountID string) ([]TransactionEntryResponseDTO, error) {
	entries, err := s.repository.GetEntriesByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	response := make([]TransactionEntryResponseDTO, 0, len(entries))
	for _, entry := range entries {
		response = append(response, TransactionEntryResponseDTO{
			ID:        entry.ID.String(),
			AccountID: entry.AccountID,
			Amount:    int64(entry.Amount.Amount),
			Currency:  string(entry.Amount.Currency),
			Direction: string(entry.Direction),
		})
	}

	return response, nil
}
