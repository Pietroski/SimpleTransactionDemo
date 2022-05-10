package transaction_controller

import (
	"database/sql"
	"net/http"

	sqlc_bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc"
	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"
	"github.com/Pietroski/SimpleTransactionDemo/internal/tools/notification"
	pkg_gin_custom_validators "github.com/Pietroski/SimpleTransactionDemo/pkg/tools/gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (c *TransactionController) getPaginatedWalletsByAccountID(
	ctx *gin.Context,
	accountID uuid.UUID,
	p *pkg_gin_custom_validators.Pagination,
) ([]sqlc_bank_account_store.Wallet, int, gin.H) {
	wallets, err := c.store.GetPaginatedWalletsByAccountID(
		ctx, sqlc_bank_account_store.GetPaginatedWalletsByAccountIDParams{
			AccountID: accountID,
			Limit:     p.Limit,
			Offset:    p.Offset,
		},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return []sqlc_bank_account_store.Wallet{},
				http.StatusNotFound,
				notification.ClientError.Response(err)
		}

		return []sqlc_bank_account_store.Wallet{},
			http.StatusInternalServerError,
			notification.ClientError.Response(err)
	}

	// TODO: beautify and filter wallets
	return wallets, 0, gin.H{}
}

func (c *TransactionController) listPaginatedEntryLogsByAccountID(
	ctx *gin.Context,
	accountID uuid.UUID,
	p *pkg_gin_custom_validators.Pagination,
) ([]sqlc_bank_account_store.EntryRecord, int, gin.H) {
	entries, err := c.store.ListPaginatedEntryLogsByAccountID(
		ctx, sqlc_bank_account_store.ListPaginatedEntryLogsByAccountIDParams{
			AccountID: accountID,
			Limit:     p.Limit,
			Offset:    p.Offset,
		},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return []sqlc_bank_account_store.EntryRecord{},
				http.StatusNotFound,
				notification.ClientError.Response(err)
		}

		return []sqlc_bank_account_store.EntryRecord{},
			http.StatusInternalServerError,
			notification.ClientError.Response(err)
	}

	// TODO: beautify and filter wallets
	return entries, 0, gin.H{}
}

func (c *TransactionController) listPaginatedCoinEntryLogsByAccountID(
	ctx *gin.Context,
	accountID uuid.UUID,
	coin manager_models.CryptoCurrencies,
	p *pkg_gin_custom_validators.Pagination,
) ([]sqlc_bank_account_store.EntryRecord, int, gin.H) {
	entries, err := c.store.ListPaginatedCoinEntryLogsByAccountID(
		ctx, sqlc_bank_account_store.ListPaginatedCoinEntryLogsByAccountIDParams{
			AccountID: accountID,
			Coin:      sqlc_bank_account_store.CryptoCurrencies(coin),
			Limit:     p.Limit,
			Offset:    p.Offset,
		},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return []sqlc_bank_account_store.EntryRecord{},
				http.StatusNotFound,
				notification.ClientError.Response(err)
		}

		return []sqlc_bank_account_store.EntryRecord{},
			http.StatusInternalServerError,
			notification.ClientError.Response(err)
	}

	// TODO: beautify and filter wallets
	return entries, 0, gin.H{}
}

func (c *TransactionController) getCoinHistory(
	ctx *gin.Context,
	accountID uuid.UUID,
	p *pkg_gin_custom_validators.Pagination,
) ([]sqlc_bank_account_store.EntryRecord, int, gin.H) {
	var coin manager_models.CoinHistoryRequest
	if err := ctx.ShouldBindQuery(&coin); err != nil {
		return []sqlc_bank_account_store.EntryRecord{},
			http.StatusInternalServerError,
			notification.ClientError.Response(err)
	}

	if !coin.Coin.IsCryptoCurrency() {
		return []sqlc_bank_account_store.EntryRecord{}, -1, nil
	}

	if ok := pkg_gin_custom_validators.IsPaginated(p); ok {
		entries, statusCode, ginResp := c.listPaginatedCoinEntryLogsByAccountID(
			ctx, accountID, coin.Coin, p,
		)
		if statusCode != 0 {
			return []sqlc_bank_account_store.EntryRecord{}, statusCode, ginResp
		}

		// TODO: beautify and filter entries
		return entries, statusCode, nil
	}

	payload := sqlc_bank_account_store.ListCoinEntryLogsByAccountIDParams{
		AccountID: accountID,
		Coin:      sqlc_bank_account_store.CryptoCurrencies(coin.Coin),
	}
	entries, err := c.store.ListCoinEntryLogsByAccountID(ctx, payload)
	if err != nil {
		return []sqlc_bank_account_store.EntryRecord{},
			http.StatusInternalServerError,
			notification.ClientError.Response(err)
	}

	// TODO: beautify and filter entries
	return entries, 0, nil
}

func (c *TransactionController) getFullHistory(
	ctx *gin.Context,
	accountID uuid.UUID,
	pagination *pkg_gin_custom_validators.Pagination,
) ([]sqlc_bank_account_store.EntryRecord, int, gin.H) {
	if ok := pkg_gin_custom_validators.IsPaginated(pagination); ok {
		entries, statusCode, ginResp := c.listPaginatedEntryLogsByAccountID(ctx, accountID, pagination)
		if statusCode != 0 {
			return []sqlc_bank_account_store.EntryRecord{}, statusCode, ginResp
		}

		// TODO: beautify and filter entries
		ctx.JSON(statusCode, entries)
		return entries, statusCode, ginResp
	}

	entries, err := c.store.ListEntryLogsByAccountID(ctx, accountID)
	if err != nil {
		return []sqlc_bank_account_store.EntryRecord{},
			http.StatusInternalServerError,
			notification.ClientError.Response(err)
	}

	// TODO: beautify and filter entries
	return entries, 0, gin.H{}
}
