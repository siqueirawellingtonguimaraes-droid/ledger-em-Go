package accounts

type AccountRepository interface {
	FindByID(id string) (*Account, error)
	FindAll() ([]Account, error)
	Save(account *Account) error
}
