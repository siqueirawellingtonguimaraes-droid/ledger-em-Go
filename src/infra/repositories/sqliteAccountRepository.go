package repositories

import (
	"database/sql"
	"ledger/src/modules/accounts"
)

type SQLiteAccountRepository struct {
	db *sql.DB
}

func NewAccountSQLiteRepository(db *sql.DB) *SQLiteAccountRepository {
	return &SQLiteAccountRepository{
		db: db,
	}
}

func (r *SQLiteAccountRepository) Save(account *accounts.Account) error {

	statements, err := r.db.Prepare("INSERT INTO accounts (id, name, currency) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer statements.Close()

	if _, err := statements.Exec(account.ID, account.Name, account.Currency); err != nil {
		return err
	}

	return err
}

func (r *SQLiteAccountRepository) FindByID(id string) (*accounts.Account, error) {
	statements, err := r.db.Prepare("SELECT * FROM accounts WHERE id = ?")
	if err != nil {
		return &accounts.Account{}, err
	}
	defer statements.Close()

	row := statements.QueryRow(id)

	var account accounts.Account

	if err := row.Scan(&account.ID, &account.Name, &account.Currency); err != nil {
		return &accounts.Account{}, err
	}

	return &account, nil
}
