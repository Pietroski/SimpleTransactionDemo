package sqlc_bank_account_store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	DepositTxParams struct {
		ToAccountID uuid.UUID        `json:"toAccountId"`
		ToWalletID  uuid.UUID        `json:"toWalletID"`
		Amount      int64            `json:"amount"`
		Coin        CryptoCurrencies `json:"coin"`
	}

	// DepositTxResult is the result of the deposit
	DepositTxResult struct {
		ToEntry           EntryRecord      `json:"toEntry"`
		ToWallet          Wallet           `json:"toWallet"`
		TransferredAmount int64            `json:"amount"`
		TransferredCoin   CryptoCurrencies `json:"coin"`
	}
)

func (store *transactionStore) DepositTx(
	ctx context.Context,
	args DepositTxParams,
) (DepositTxResult, error) {
	var result DepositTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.ToEntry, err = q.LogEntry(
			ctx,
			LogEntryParams{
				AccountID: args.ToAccountID,
				WalletID:  args.ToWalletID,
				Coin:      args.Coin,
				Amount:    args.Amount,
				CreatedAt: time.Now(),
			},
		)
		if err != nil {
			return err
		}

		result.ToWallet, err = store.UpdateAccountWalletBalance(
			ctx,
			UpdateAccountWalletBalanceParams{
				AccountID: args.ToAccountID,
				WalletID:  args.ToWalletID,
				Amount:    args.Amount,
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
