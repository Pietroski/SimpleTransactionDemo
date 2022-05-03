package transaction_factory

import (
	"github.com/gin-gonic/gin"

	bank_account_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/manager/transactions"
)

type TransactionFactory struct {
	transactionController *bank_account_controller.TransactionController
}

func NewTransactionFactory(controller *bank_account_controller.TransactionController) *TransactionFactory {
	// TODO: apply validations for arguments if needed

	factory := &TransactionFactory{
		transactionController: controller,
	}

	return factory
}

func (f *TransactionFactory) Handle(engine *gin.RouterGroup) *gin.RouterGroup {
	transactionGroup := engine.Group("/transactions")
	transactionGroup.Use()
	{
		transactionGroup.POST("/deposit", f.transactionController.Deposit)
		transactionGroup.POST("/withdraw", f.transactionController.Withdraw)
		transactionGroup.GET("/balance", f.transactionController.GetBalance)
		transactionGroup.GET("/history", f.transactionController.GetHistory)
	}

	return transactionGroup
}
