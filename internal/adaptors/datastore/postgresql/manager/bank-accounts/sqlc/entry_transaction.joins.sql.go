// Code generated by go-migrate. DO NOT EDIT.
// source: entry_transaction.joins.sql

package sqlc_bank_account_store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const logTransactionWithEntriesSimplified = `-- name: LogTransactionWithEntriesSimplified :one


WITH first_insert AS (
    INSERT INTO entry_record
        (account_id, wallet_id, coin, amount, created_at)
        VALUES ($1, $2, $5, $6,
                $7)
        RETURNING row_id, account_id, wallet_id, coin, amount, created_at),
     second_insert AS (
         INSERT INTO entry_record
             (account_id, wallet_id, coin, amount, created_at)
             VALUES ($3, $4, $5, $6,
                     $7)
             RETURNING row_id, account_id, wallet_id, coin, amount, created_at)
INSERT
INTO transaction_record
(from_account_id, from_wallet_id, to_account_id, to_wallet_id, coin, amount, created_at)
VALUES ($1, $2, $3, $4,
        $5, $6,
        $7)
RETURNING row_id, from_account_id, from_wallet_id, to_account_id, to_wallet_id, coin, amount, created_at
`

type LogTransactionWithEntriesSimplifiedParams struct {
	FromAccountID uuid.UUID        `json:"fromAccountID"`
	FromWalletID  uuid.UUID        `json:"fromWalletID"`
	ToAccountID   uuid.UUID        `json:"toAccountID"`
	ToWalletID    uuid.UUID        `json:"toWalletID"`
	Coin          CryptoCurrencies `json:"coin"`
	Amount        int64            `json:"amount"`
	CreatedAt     time.Time        `json:"createdAt"`
}

// -- name: LogTransactionWithEntries :one
// WITH transaction_record
//          (from_account_id, from_wallet_id, to_account_id, to_wallet_id, coin, amount, created_at)
//          AS (VALUES (go-migrate.arg(from_account_id), go-migrate.arg(from_wallet_id), go-migrate.arg(to_account_id),
//                      go-migrate.arg(to_wallet_id),
//                      go-migrate.arg(coin), go-migrate.arg(amount),
//                      go-migrate.arg(created_at))),
//      ins1 AS (
//          INSERT INTO entry_record
//              (account_id, wallet_id, coin, amount, created_at)
//              SELECT (from_account_id, from_wallet_id, coin, amount, created_at)
//              FROM transaction_record),
//      ins2 AS (
//          INSERT INTO entry_record
//              (account_id, wallet_id, coin, amount, created_at)
//              SELECT (to_account_id, to_wallet_id, coin, amount, created_at))
// SELECT * FROM transaction_record;
//
//
//
//
func (q *Queries) LogTransactionWithEntriesSimplified(ctx context.Context, arg LogTransactionWithEntriesSimplifiedParams) (TransactionRecord, error) {
	row := q.queryRow(ctx, q.logTransactionWithEntriesSimplifiedStmt, logTransactionWithEntriesSimplified,
		arg.FromAccountID,
		arg.FromWalletID,
		arg.ToAccountID,
		arg.ToWalletID,
		arg.Coin,
		arg.Amount,
		arg.CreatedAt,
	)
	var i TransactionRecord
	err := row.Scan(
		&i.RowID,
		&i.FromAccountID,
		&i.FromWalletID,
		&i.ToAccountID,
		&i.ToWalletID,
		&i.Coin,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
