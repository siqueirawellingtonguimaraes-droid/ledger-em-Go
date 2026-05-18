package transactions

import "github.com/gin-gonic/gin"

func SetupTransactionRouters(r *gin.Engine, service *TransactionService) {
	transactionGroup := r.Group("/transactions")
	{
		transactionGroup.POST("/", func(ctx *gin.Context) {
			var body TransactionCreateDTO

			if err := ctx.ShouldBindJSON(&body); err != nil {
				ctx.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err := service.Create(&body); err != nil {
				ctx.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			ctx.Status(201)
		})

		transactionGroup.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			response, err := service.GetByID(id)
			if err != nil {
				ctx.JSON(404, gin.H{
					"error": err.Error(),
				})
				return
			}

			ctx.JSON(200, response)
		})

		transactionGroup.GET("/", func(ctx *gin.Context) {
			response, err := service.GetAll()
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			ctx.JSON(200, response)
		})

		transactionGroup.GET("/entries/:accountID", func(ctx *gin.Context) {
			accountID := ctx.Param("accountID")

			response, err := service.GetEntriesByAccountID(accountID)
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			ctx.JSON(200, response)
		})
	}
}
