package transaction_factory

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	bank_account_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/manager/transactions"
	mocked_auth_middleware "github.com/Pietroski/SimpleTransactionDemo/internal/middlewares/auth/mocked"
	internal_gin_custom_validators "github.com/Pietroski/SimpleTransactionDemo/internal/tools/gin/validators/manager/transactions"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// TODO: handle Error
		_ = v.RegisterValidation(
			"CoinCustomValidation",
			internal_gin_custom_validators.CoinCustomValidation,
		)
	}

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
