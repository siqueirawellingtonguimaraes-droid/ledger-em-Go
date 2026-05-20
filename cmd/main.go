package main

import (
	"ledger/src/config"
	"ledger/src/infra/db"
	"ledger/src/infra/repositories"
	"ledger/src/modules/accounts"
	"ledger/src/modules/transactions"
	"log"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5500",
			"http://127.0.0.1:5500",
		},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	config.SetupRoutes(r, accountModule, transactionModule)

	log.Fatal(r.Run(":8080"))
}
