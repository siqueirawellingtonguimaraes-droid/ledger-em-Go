package transactions

type TransactionRepository interface {
	Save(transaction *Transaction) error
	FindByID(id string) (Transaction, error)
	FindAll() ([]Transaction, error)
	GetEntriesByAccountID(accountID string) ([]TransactionEntry, error)
}
