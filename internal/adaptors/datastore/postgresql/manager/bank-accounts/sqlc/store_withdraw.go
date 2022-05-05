package sqlc_bank_account_store

import (
	"context"

	"github.com/google/uuid"
)

type (
	WithdrawTxParams struct {
		FromAccountID uuid.UUID        `json:"fromAccountId"`
		FromWalletID  uuid.UUID        `json:"fromWalletID"`
		Amount        int64            `json:"amount"`
		Coin          CryptoCurrencies `json:"coin"`
	}
)

func (store *transactionStore) WithdrawTx(
	ctx context.Context,
	args WithdrawTxParams,
) (Wallet, error) {
	return store.UpdateFromWallet(ctx, TransferTxParams{
		FromAccountID: args.FromAccountID,
		FromWalletID:  args.FromWalletID,
		Amount:        args.Amount,
		Coin:          args.Coin,
	})
}
