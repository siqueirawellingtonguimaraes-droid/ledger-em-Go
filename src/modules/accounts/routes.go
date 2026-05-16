package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAccountRouters(r *gin.Engine, service *AccountService) {
	accountGroup := r.Group("/accounts")
	{
		accountGroup.POST("/", func(ctx *gin.Context) {
			var body AccountCreateDTO

			if err := ctx.ShouldBindJSON(&body); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err,
				})
			}

			err := service.Create(&body)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			}

			ctx.JSON(http.StatusCreated, gin.H{
				"message": "Conta criada com sucesso",
			})
		})

		accountGroup.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Get account by ID endpoint",
			})
		})
	}
}
