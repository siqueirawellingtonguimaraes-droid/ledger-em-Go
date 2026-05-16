package main

import (
	"ledger/src/config"
	"ledger/src/infra/db"
	"ledger/src/infra/repositories"
	"ledger/src/modules/accounts"
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

	repo := repositories.NewAccountSQLiteRepository(db)

	accoutService := accounts.NewAccountService(repo)

	r := gin.Default()

	config.SetupRoutes(r, accoutService)

	log.Fatal(r.Run(":8080"))
}
