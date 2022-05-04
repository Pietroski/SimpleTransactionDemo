package sqlc_bank_account_store

import (
	"context"
	"github.com/google/uuid"
)

type (
	DepositTxParams struct {
		ToAccountID uuid.UUID        `json:"toAccountId"`
		ToWalletID  uuid.UUID        `json:"toWalletID"`
		Amount      int64            `json:"amount"`
		Coin        CryptoCurrencies `json:"coin"`
	}
)

func (store *transactionStore) DepositTx(
	ctx context.Context,
	args DepositTxParams,
) (Wallet, error) {
	return store.UpdateToWallet(ctx, TransferTxParams{
		ToAccountID: args.ToAccountID,
		ToWalletID:  args.ToWalletID,
		Amount:      args.Amount,
		Coin:        args.Coin,
	})
}
