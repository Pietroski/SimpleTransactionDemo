package transaction_controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"

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
	authInfo, err := mocked_auth_middleware.MockedAuthMiddlewareExtractor(ctx)
	if err != nil {
		if errors.As(err, &mocked_auth_middleware.ErrToAssertVar) {
			ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
			return
		}

		// errors.As(err, &pkg_auth.ErrInvalidAuthBearToken)
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	var transferPayload manager_models.TransactionRequest
	if err = ctx.ShouldBindJSON(&transferPayload); err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	rawAccountID := authInfo.AccountID
	accountID, err := uuid.Parse(rawAccountID.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	fromWallet, err := c.store.GetTxWallet(
		ctx,
		sqlc_bank_account_store.GetTxWalletParams{
			AccountID: accountID,
			Coin:      sqlc_bank_account_store.CryptoCurrencies(transferPayload.Coin),
		},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, notification.ClientError.Response(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	toWallet, err := c.store.GetTxWallet(
		ctx,
		sqlc_bank_account_store.GetTxWalletParams{
			AccountID: transferPayload.ToAccountID,
			Coin:      sqlc_bank_account_store.CryptoCurrencies(transferPayload.Coin),
		},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, notification.ClientError.Response(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	txResult, err := c.store.TransferTx(
		ctx,
		sqlc_bank_account_store.TransferTxParams{
			FromAccountID: fromWallet.AccountID,
			FromWalletID:  fromWallet.WalletID,
			ToAccountID:   toWallet.AccountID,
			ToWalletID:    toWallet.WalletID,
			Amount:        transferPayload.Amount,
			Coin:          sqlc_bank_account_store.CryptoCurrencies(transferPayload.Coin),
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	// TODO: beautify and filter txResult
	ctx.JSON(http.StatusOK, txResult)
}

func (c *TransactionController) Deposit(ctx *gin.Context) {
	authInfo, err := mocked_auth_middleware.MockedAuthMiddlewareExtractor(ctx)
	if err != nil {
		if errors.As(err, &mocked_auth_middleware.ErrToAssertVar) {
			ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
			return
		}

		// errors.As(err, &pkg_auth.ErrInvalidAuthBearToken)
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	var depositPayload manager_models.DepositRequest
	if err = ctx.ShouldBindJSON(&depositPayload); err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	rawAccountID := authInfo.AccountID
	accountID, err := uuid.Parse(rawAccountID.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	wallet, err := c.store.GetTxWallet(
		ctx,
		sqlc_bank_account_store.GetTxWalletParams{
			AccountID: accountID,
			Coin:      sqlc_bank_account_store.CryptoCurrencies(depositPayload.Coin),
		},
	)

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
	authInfo, err := mocked_auth_middleware.MockedAuthMiddlewareExtractor(ctx)
	if err != nil {
		if errors.As(err, &mocked_auth_middleware.ErrToAssertVar) {
			ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
			return
		}

		// errors.As(err, &pkg_auth.ErrInvalidAuthBearToken)
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	var withdrawPayload manager_models.WithdrawRequest
	if err = ctx.ShouldBindJSON(&withdrawPayload); err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	rawAccountID := authInfo.AccountID
	accountID, err := uuid.Parse(rawAccountID.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	wallet, err := c.store.GetTxWallet(
		ctx,
		sqlc_bank_account_store.GetTxWalletParams{
			AccountID: accountID,
			Coin:      sqlc_bank_account_store.CryptoCurrencies(withdrawPayload.Coin),
		},
	)

	withdrawResult, err := c.store.WithdrawTx(ctx, sqlc_bank_account_store.WithdrawTxParams{
		FromAccountID: wallet.AccountID,
		FromWalletID:  wallet.WalletID,
		Amount:        withdrawPayload.Amount,
		Coin:          sqlc_bank_account_store.CryptoCurrencies(withdrawPayload.Coin),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	// TODO: beautify and filter withdrawResult
	ctx.JSON(http.StatusOK, withdrawResult)
}

func (c *TransactionController) GetBalance(ctx *gin.Context) {
	c.getWallet(ctx)
}

func (c *TransactionController) getWallet(ctx *gin.Context) {
	// TODO: implement me!!
	authInfo, err := mocked_auth_middleware.MockedAuthMiddlewareExtractor(ctx)
	if err != nil {
		if errors.As(err, &mocked_auth_middleware.ErrToAssertVar) {
			ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
			return
		}

		// errors.As(err, &pkg_auth.ErrInvalidAuthBearToken)
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	var balancePayload manager_models.BalanceRequest
	if err = ctx.ShouldBindJSON(&balancePayload); err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	rawAccountID := authInfo.AccountID
	accountID, err := uuid.Parse(rawAccountID.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
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
	// TODO: implement me!!
	authInfo, err := mocked_auth_middleware.MockedAuthMiddlewareExtractor(ctx)
	if err != nil {
		if errors.As(err, &mocked_auth_middleware.ErrToAssertVar) {
			ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
			return
		}

		// errors.As(err, &pkg_auth.ErrInvalidAuthBearToken)
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	rawAccountID := authInfo.AccountID
	accountID, err := uuid.Parse(rawAccountID.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
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
	// TODO: implement me!!
	authInfo, err := mocked_auth_middleware.MockedAuthMiddlewareExtractor(ctx)
	if err != nil {
		if errors.As(err, &mocked_auth_middleware.ErrToAssertVar) {
			ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
			return
		}

		// errors.As(err, &pkg_auth.ErrInvalidAuthBearToken)
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	rawAccountID := authInfo.AccountID
	accountID, err := uuid.Parse(rawAccountID.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
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
