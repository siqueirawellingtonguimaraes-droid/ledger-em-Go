package transactions

import "github.com/gin-gonic/gin"

type Module struct {
	Service *TransactionService
}

func NewModule(service *TransactionService) *Module {
	return &Module{Service: service}
}

func (m *Module) Register(r *gin.Engine) {
	SetupTransactionRouters(r, m.Service)
}
