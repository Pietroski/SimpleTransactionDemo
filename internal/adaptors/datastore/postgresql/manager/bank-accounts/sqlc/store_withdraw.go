package sqlc_bank_account_store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	WithdrawTxParams struct {
		FromAccountID uuid.UUID        `json:"fromAccountId"`
		FromWalletID  uuid.UUID        `json:"fromWalletID"`
		Amount        int64            `json:"amount"`
		Coin          CryptoCurrencies `json:"coin"`
	}

	// WithdrawTxResult is the result of the withdrawal
	WithdrawTxResult struct {
		FromEntry         EntryRecord      `json:"toEntry"`
		FromWallet        Wallet           `json:"toWallet"`
		TransferredAmount int64            `json:"amount"`
		TransferredCoin   CryptoCurrencies `json:"coin"`
	}
)

func (store *transactionStore) WithdrawTx(
	ctx context.Context,
	args WithdrawTxParams,
) (WithdrawTxResult, error) {
	var result WithdrawTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.FromEntry, err = q.LogEntry(
			ctx,
			LogEntryParams{
				AccountID: args.FromAccountID,
				WalletID:  args.FromWalletID,
				Coin:      args.Coin,
				Amount:    -args.Amount,
				CreatedAt: time.Now(),
			},
		)
		if err != nil {
			return err
		}

		result.FromWallet, err = store.UpdateAccountWalletBalance(
			ctx,
			UpdateAccountWalletBalanceParams{
				AccountID: args.FromAccountID,
				WalletID:  args.FromWalletID,
				Amount:    -args.Amount,
				Coin:      args.Coin,
				UpdatedAt: time.Now(),
			},
		)

		return nil
	})

	result.TransferredAmount = args.Amount
	result.TransferredCoin = args.Coin
	return result, err
}
