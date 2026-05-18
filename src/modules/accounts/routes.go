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
				return
			}

			err := service.Create(&body)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				return
			}

			ctx.JSON(http.StatusCreated, gin.H{
				"message": "Conta criada com sucesso",
			})
		})

		accountGroup.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			account, err := service.GetByID(id)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, account)
		})
	}
}
