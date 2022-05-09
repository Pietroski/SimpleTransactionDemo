package transaction_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	sqlc_bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc"
	mocked_auth_middleware "github.com/Pietroski/SimpleTransactionDemo/internal/middlewares/auth/mocked"
	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"
	"github.com/Pietroski/SimpleTransactionDemo/internal/tools/notification"
)

type TransactionController struct {
	store sqlc_bank_account_store.Store
}

func NewTransactionController(store sqlc_bank_account_store.Store) *TransactionController {
	// TODO: apply validations for arguments if needed

	controller := &TransactionController{
		store: store,
	}

	return controller
}

func (c *TransactionController) Transfer(ctx *gin.Context) {
	accountID, statusCode, ginResp := mocked_auth_middleware.AccountIdCtxExtractor(ctx)
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	var transferPayload manager_models.TransactionRequest
	if err := ctx.ShouldBindJSON(&transferPayload); err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	fromWallet, statusCode, ginResp := c.getTxWallet(ctx, accountID, transferPayload.Coin.String())
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}
	toWallet, statusCode, ginResp := c.getTxWallet(ctx, transferPayload.ToAccountID, transferPayload.Coin.String())
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	txResult, statusCode, ginResp := c.tx(
		ctx,
		fromWallet, toWallet,
		transferPayload.Amount,
		transferPayload.Coin.String(),
	)
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	// TODO: beautify and filter txResult
	ctx.JSON(http.StatusOK, txResult)
}

func (c *TransactionController) Deposit(ctx *gin.Context) {
	accountID, statusCode, ginResp := mocked_auth_middleware.AccountIdCtxExtractor(ctx)
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	var depositPayload manager_models.DepositRequest
	if err := ctx.ShouldBindJSON(&depositPayload); err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	wallet, statusCode, ginResp := c.getTxWallet(ctx, accountID, depositPayload.Coin.String())
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	depositResult, err := c.store.DepositTx(ctx, sqlc_bank_account_store.DepositTxParams{
		ToAccountID: wallet.AccountID,
		ToWalletID:  wallet.WalletID,
		Amount:      depositPayload.Amount,
		Coin:        sqlc_bank_account_store.CryptoCurrencies(depositPayload.Coin),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	// TODO: beautify and filter depositResult
	ctx.JSON(http.StatusOK, depositResult)
}

func (c *TransactionController) Withdraw(ctx *gin.Context) {
	accountID, statusCode, ginResp := mocked_auth_middleware.AccountIdCtxExtractor(ctx)
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	var depositPayload manager_models.WithdrawRequest
	if err := ctx.ShouldBindJSON(&depositPayload); err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	wallet, statusCode, ginResp := c.getTxWallet(ctx, accountID, depositPayload.Coin.String())
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	withdrawResult, err := c.store.WithdrawTx(ctx, sqlc_bank_account_store.WithdrawTxParams{
		FromAccountID: wallet.AccountID,
		FromWalletID:  wallet.WalletID,
		Amount:        depositPayload.Amount,
		Coin:          sqlc_bank_account_store.CryptoCurrencies(depositPayload.Coin),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	// TODO: beautify and filter depositResult
	ctx.JSON(http.StatusOK, withdrawResult)
}

func (c *TransactionController) GetBalance(ctx *gin.Context) {
	c.getWallet(ctx)
}

func (c *TransactionController) getWallet(ctx *gin.Context) {
	accountID, statusCode, ginResp := mocked_auth_middleware.AccountIdCtxExtractor(ctx)
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	var balancePayload manager_models.BalanceRequest
	if err := ctx.ShouldBindJSON(&balancePayload); err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	wallet, err := c.store.GetTxWallet(ctx, sqlc_bank_account_store.GetTxWalletParams{
		AccountID: accountID,
		Coin:      sqlc_bank_account_store.CryptoCurrencies(balancePayload.Coin),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	// TODO: beautify and filter wallet
	ctx.JSON(http.StatusOK, wallet)

}

func (c *TransactionController) GetWallets(ctx *gin.Context) {
	accountID, statusCode, ginResp := mocked_auth_middleware.AccountIdCtxExtractor(ctx)
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	wallets, err := c.store.GetAccountWallets(ctx, accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	// TODO: beautify and filter wallets
	ctx.JSON(http.StatusOK, wallets)
}

func (c *TransactionController) GetHistory(ctx *gin.Context) {
	accountID, statusCode, ginResp := mocked_auth_middleware.AccountIdCtxExtractor(ctx)
	if statusCode != 0 {
		ctx.JSON(statusCode, ginResp)
		return
	}

	entries, err := c.store.ListEntryLogsByAccountID(ctx, accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	// TODO: beautify and filter entries
	ctx.JSON(http.StatusOK, entries)
}
