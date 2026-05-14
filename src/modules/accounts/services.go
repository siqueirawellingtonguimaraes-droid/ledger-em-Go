package accounts

type accountService struct {
	repository AccountRepository
}

func NewAccountService(repository AccountRepository) *accountService {
	return &accountService{
		repository: repository,
	}
}

func (s *accountService) Create(dto *AccountCreateDTO) error {
	account := Account{
		Name: dto.Name,
	}
	return s.repository.Save(&account)
}

func (s *accountService) GetByID(id string) (*AccountResponseDTO, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &AccountResponseDTO{
		ID:   account.ID,
		Name: account.Name,
	}, nil
}
