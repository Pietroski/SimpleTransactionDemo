package transaction_controller

import (
	"database/sql"
	sqlc_bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc"
	"github.com/Pietroski/SimpleTransactionDemo/internal/tools/notification"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (c *TransactionController) getTxWallet(
	ctx *gin.Context,
	accountID uuid.UUID,
	coin string,
) (sqlc_bank_account_store.Wallet, int, gin.H) {
	payload := sqlc_bank_account_store.GetTxWalletParams{
		AccountID: accountID,
		Coin:      sqlc_bank_account_store.CryptoCurrencies(coin),
	}
	wallet, err := c.store.GetTxWallet(ctx, payload)
	if err != nil {
		if err == sql.ErrNoRows {

			return sqlc_bank_account_store.Wallet{},
				http.StatusNotFound,
				notification.ClientError.Response(err)
		}

		return sqlc_bank_account_store.Wallet{},
			http.StatusInternalServerError,
			notification.ClientError.Response(err)
	}

	return wallet, 0, gin.H{}
}

func (c *TransactionController) tx(
	ctx *gin.Context,
	fromWallet, toWallet sqlc_bank_account_store.Wallet,
	amount int64,
	coin string,
) (sqlc_bank_account_store.TransferTxResult, int, gin.H) {
	txResult, err := c.store.TransferTx(
		ctx,
		sqlc_bank_account_store.TransferTxParams{
			FromAccountID: fromWallet.AccountID,
			FromWalletID:  fromWallet.WalletID,
			ToAccountID:   toWallet.AccountID,
			ToWalletID:    toWallet.WalletID,
			Amount:        amount,
			Coin:          sqlc_bank_account_store.CryptoCurrencies(coin),
		},
	)
	if err != nil {
		return sqlc_bank_account_store.TransferTxResult{},
			http.StatusInternalServerError,
			notification.ClientError.Response(err)
	}

	return txResult, 0, gin.H{}
}
