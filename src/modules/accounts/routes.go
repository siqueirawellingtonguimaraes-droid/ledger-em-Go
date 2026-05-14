package accounts

import "github.com/gin-gonic/gin"

func SetupAccountRouters(r *gin.Engine) {
	accountGroup := r.Group("/accounts")
	{
		accountGroup.POST("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Create account endpoint",
			})
		})

		accountGroup.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Get account by ID endpoint",
			})
		})
	}
}
