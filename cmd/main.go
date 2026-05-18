package main

import (
	"ledger/src/config"
	"ledger/src/infra/db"
	"ledger/src/infra/repositories"
	"ledger/src/modules/accounts"
	"ledger/src/modules/transactions"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := config.InitDB(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := db.NewSQLiteConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	transactionRepo := repositories.NewTransactionSQLiteRepository(db)
	transactionModule := transactions.NewModule(transactions.NewTransactionService(transactionRepo))

	accountrepo := repositories.NewAccountSQLiteRepository(db)
	accountModule := accounts.NewModule(accounts.NewAccountService(accountrepo, transactionRepo))

	r := gin.Default()

	config.SetupRoutes(r, accountModule, transactionModule)

	log.Fatal(r.Run(":8080"))
}
