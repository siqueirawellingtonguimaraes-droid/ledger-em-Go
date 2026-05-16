package accounts

type AccountRepository interface {
	FindByID(id string) (*Account, error)
	Save(account *Account) error
}
