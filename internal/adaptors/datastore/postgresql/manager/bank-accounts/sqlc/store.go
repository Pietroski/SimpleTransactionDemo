package sqlc_bank_account_store

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries
type Store interface {
	Querier
	TransferTx(
		ctx context.Context,
		args TransferTxParams,
	) (TransferTxResult, error)
	DepositTx(
		ctx context.Context,
		args DepositTxParams,
	) (Wallet, error)
	WithdrawTx(
		ctx context.Context,
		args WithdrawTxParams,
	) (Wallet, error)
}

type transactionStore struct {
	*Queries
	db *sql.DB
}

// NewStore instantiates a devices store object returning the store interface.
func NewStore(db *sql.DB) Store {
	store := &transactionStore{
		Queries: New(db),
		db:      db,
	}

	return store
}

// ExecTx executes a function within a database transaction
func (store *transactionStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//q := New(tx)
	q := store.Queries
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
