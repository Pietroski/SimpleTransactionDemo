package transaction_factory

import (
	"github.com/gin-gonic/gin"

	bank_account_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/manager/transactions"
	mocked_auth_middleware "github.com/Pietroski/SimpleTransactionDemo/internal/middlewares/auth/mocked"
)

type TransactionFactory struct {
	transactionController *bank_account_controller.TransactionController
}

func NewTransactionFactory(
	controller *bank_account_controller.TransactionController,
) *TransactionFactory {
	// TODO: apply validations for arguments if needed

	factory := &TransactionFactory{
		transactionController: controller,
	}

	return factory
}

func (f *TransactionFactory) Handle(engine *gin.RouterGroup) *gin.RouterGroup {
	transactionGroup := engine.Group("/transactions")
	transactionGroup.Use(mocked_auth_middleware.MockedAuthMiddleware)
	{
		transactionGroup.POST("/transfer", f.transactionController.Transfer)
		transactionGroup.POST("/deposit", f.transactionController.Deposit)
		transactionGroup.POST("/withdraw", f.transactionController.Withdraw)
		transactionGroup.GET("/balance", f.transactionController.GetWalletBalance)
		transactionGroup.GET("/wallets", f.transactionController.GetWallets)
		transactionGroup.GET("/history", f.transactionController.GetHistory)
	}

	return transactionGroup
}
