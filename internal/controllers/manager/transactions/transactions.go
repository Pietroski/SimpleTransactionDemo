package transaction_controller

import (
	bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/manager/bank-accounts"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	store bank_account_store.Store
}

func NewTransactionController(store bank_account_store.Store) *TransactionController {
	// TODO: apply validations for arguments if needed

	controller := &TransactionController{
		store: store,
	}

	return controller
}

func (c *TransactionController) Deposit(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *TransactionController) Withdraw(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *TransactionController) GetBalance(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *TransactionController) GetHistory(ctx *gin.Context) {
	// TODO: implement me!!
}
