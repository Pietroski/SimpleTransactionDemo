package sqlc_bank_account_store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TransferTxParams contains the input parameters of the transfer transaction
type (
	TransferTxParams struct {
		FromAccountID uuid.UUID        `json:"fromAccountId"`
		FromWalletID  uuid.UUID        `json:"fromWalletID"`
		ToAccountID   uuid.UUID        `json:"toAccountId"`
		ToWalletID    uuid.UUID        `json:"toWalletID"`
		Amount        int64            `json:"amount"`
		Coin          CryptoCurrencies `json:"coin"`
	}

	// TransferTxResult is the result of the transfer transaction
	TransferTxResult struct {
		TransactionRecord TransactionRecord `json:"transactionRecord"`
		FromEntry         EntryRecord       `json:"fromEntry"`
		ToEntry           EntryRecord       `json:"toEntry"`
		FromWallet        Wallet            `json:"fromWallet"`
		ToWallet          Wallet            `json:"toWallet"`
		TransferredAmount int64             `json:"amount"`
		TransferredCoin   CryptoCurrencies  `json:"coin"`
	}
)

// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, add account entries, and update accounts' balance within a database transaction
func (store *transactionStore) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.TransactionRecord, err = q.LogTransaction(
			ctx,
			LogTransactionParams{
				FromAccountID: args.FromAccountID,
				FromWalletID:  args.FromWalletID,
				ToAccountID:   args.ToAccountID,
				ToWalletID:    args.ToWalletID,
				Coin:          args.Coin,
				Amount:        args.Amount,
				CreatedAt:     time.Now(),
			},
		)
		if err != nil {
			return err
		}

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

		if args.FromAccountID.ID() < args.ToAccountID.ID() {
			result.FromWallet, result.ToWallet, err = store.UpdateWallets(
				ctx, args, FromAccountFirst,
			)
		} else {
			result.FromWallet, result.ToWallet, err = store.UpdateWallets(
				ctx, args, ToAccountFirst,
			)
		}

		return err
	})

	result.TransferredAmount = args.Amount
	result.TransferredCoin = args.Coin
	return result, err
}

type TxOrder int

const (
	FromAccountFirst TxOrder = iota
	ToAccountFirst
)

func (store *transactionStore) UpdateWallets(
	ctx context.Context,
	args TransferTxParams,
	txOrder TxOrder,
) (fromWallet Wallet, toWallet Wallet, err error) {
	switch txOrder {
	case FromAccountFirst:
		if fromWallet, err = store.UpdateFromWallet(ctx, args); err != nil {
			return
		}

		toWallet, err = store.UpdateToWallet(ctx, args)
	case ToAccountFirst:
		if toWallet, err = store.UpdateToWallet(ctx, args); err != nil {
			return
		}

		fromWallet, err = store.UpdateFromWallet(ctx, args)
	}

	return
}

func (store *transactionStore) UpdateFromWallet(
	ctx context.Context,
	args TransferTxParams,
) (Wallet, error) {
	return store.UpdateAccountWalletBalance(
		ctx,
		UpdateAccountWalletBalanceParams{
			AccountID: args.FromAccountID,
			WalletID:  args.FromWalletID,
			Amount:    -args.Amount,
			Coin:      args.Coin,
			UpdatedAt: time.Now(),
		},
	)
}

func (store *transactionStore) UpdateToWallet(
	ctx context.Context,
	args TransferTxParams,
) (Wallet, error) {
	return store.UpdateAccountWalletBalance(
		ctx,
		UpdateAccountWalletBalanceParams{
			AccountID: args.ToAccountID,
			WalletID:  args.ToWalletID,
			Amount:    args.Amount,
			Coin:      args.Coin,
			UpdatedAt: time.Now(),
		},
	)
}
