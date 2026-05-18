package config

import (
	"github.com/gin-gonic/gin"
)

type Module interface {
	Register(r *gin.Engine)
}

func SetupRoutes(r *gin.Engine, modules ...Module) {
	for _, module := range modules {
		module.Register(r)
	}
}
