package main

import (
	"ledger/src/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.SetupRoutes(r)

	r.Run()
}
