package accounts

import "github.com/gin-gonic/gin"

type Module struct {
	Service *AccountService
}

func NewModule(service *AccountService) *Module {
	return &Module{Service: service}
}

func (m *Module) Register(r *gin.Engine) {
	SetupAccountRouters(r, m.Service)
}
