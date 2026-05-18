package repositories

import (
	"database/sql"
	vo "ledger/src/VO"
	"ledger/src/modules/transactions"

	"github.com/google/uuid"
)

type SQLiteTransactionRepository struct {
	db *sql.DB
}

func NewTransactionSQLiteRepository(db *sql.DB) *SQLiteTransactionRepository {
	return &SQLiteTransactionRepository{
		db: db,
	}
}

func (r *SQLiteTransactionRepository) Save(transaction *transactions.Transaction) error {

	statements, err := r.db.Prepare("INSERT INTO transactions (id, description) VALUES (?, ?)")
	if err != nil {
		return err
	}

	defer statements.Close()
	_, err = statements.Exec(transaction.ID, transaction.Description)
	if err != nil {
		return err
	}

	statements, err = r.db.Prepare("INSERT INTO transaction_entries (id, transaction_id, account_id, amount, currency, direction) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer statements.Close()
	for _, entry := range transaction.Entries {
		_, err = statements.Exec(entry.ID.String(), transaction.ID.String(), entry.AccountID, entry.Amount.Amount, entry.Amount.Currency, entry.Direction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SQLiteTransactionRepository) FindByID(id string) (*transactions.Transaction, error) {
	statements, err := r.db.Prepare("SELECT * FROM transactions WHERE id=?")
	if err != nil {
		return &transactions.Transaction{}, err
	}
	defer statements.Close()
	var transaction transactions.Transaction

	row := statements.QueryRow(id)
	if err = row.Scan(&transaction.ID, &transaction.Description); err != nil {
		return &transactions.Transaction{}, err
	}

	statements, err = r.db.Prepare("SELECT * FROM transaction_entries WHERE transaction_id=?")
	if err != nil {
		return &transactions.Transaction{}, err
	}

	rows, err := statements.Query(id)
	if err != nil {
		return &transactions.Transaction{}, err
	}

	defer statements.Close()
	for rows.Next() {
		var (
			entryID   string
			accountID string
			amount    uint64
			currency  string
			direction string
		)

		if err = rows.Scan(&entryID, &transaction.ID, &accountID, &amount, &currency, &direction); err != nil {
			return &transactions.Transaction{}, err
		}

		money, err := vo.NewMoney(amount, vo.Currency(currency))
		if err != nil {
			return &transactions.Transaction{}, err
		}

		entry := transactions.TransactionEntry{
			ID:        uuid.MustParse(entryID),
			AccountID: accountID,
			Amount:    money,
			Direction: vo.Direction(direction),
		}
		transaction.Entries = append(transaction.Entries, entry)
	}

	return &transaction, nil
}

func (r *SQLiteTransactionRepository) FindAll() ([]transactions.Transaction, error) {
	statements, err := r.db.Prepare("SELECT * FROM transactions")
	if err != nil {
		return []transactions.Transaction{}, err
	}
	defer statements.Close()

	rows, err := statements.Query()
	if err != nil {
		return []transactions.Transaction{}, err
	}
	var transactionsList []transactions.Transaction
	for rows.Next() {
		var transaction transactions.Transaction
		if err = rows.Scan(&transaction.ID, &transaction.Description); err != nil {
			return []transactions.Transaction{}, err
		}
		transactionsList = append(transactionsList, transaction)
	}
	return transactionsList, nil
}

func (r *SQLiteTransactionRepository) GetEntriesByAccountID(accountID string) ([]transactions.TransactionEntry, error) {
	statements, err := r.db.Prepare("SELECT * FROM transaction_entries WHERE account_id=?")
	if err != nil {
		return []transactions.TransactionEntry{}, err
	}
	defer statements.Close()

	rows, err := statements.Query(accountID)
	if err != nil {
		return []transactions.TransactionEntry{}, err
	}

	var entries []transactions.TransactionEntry
	for rows.Next() {
		var (
			entryID       string
			transactionID string
			accountID     string
			amount        uint64
			currency      string
			direction     string
		)

		if err = rows.Scan(&entryID, &transactionID, &accountID, &amount, &currency, &direction); err != nil {
			return []transactions.TransactionEntry{}, err
		}

		money, err := vo.NewMoney(amount, vo.Currency(currency))
		if err != nil {
			return []transactions.TransactionEntry{}, err
		}

		entry := transactions.TransactionEntry{
			ID:        uuid.MustParse(entryID),
			AccountID: accountID,
			Amount:    money,
			Direction: vo.Direction(direction),
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
