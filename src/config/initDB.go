package config

import (
	"ledger/src/infra"
	"os"
)

func InitDB() error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlBytes, err := os.ReadFile("sql/db.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return err
	}

	return nil
}
