package config

import (
	"ledger/src/modules/accounts"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, service *accounts.AccountService) {
	accounts.SetupAccountRouters(r, service)
}
