// Code generated by sqlc. DO NOT EDIT.
// source: logs.transactions.sql

package bank_account_store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const listFromUserTransactionLogs = `-- name: ListFromUserTransactionLogs :many
SELECT row_id, from_user_id, from_wallet_id, to_user_id, to_wallet_id, coin, amount, created_at
FROM transaction_record
WHERE from_user_id = $1
ORDER BY row_id
`

func (q *Queries) ListFromUserTransactionLogs(ctx context.Context, fromUserID uuid.UUID) ([]TransactionRecord, error) {
	rows, err := q.query(ctx, q.listFromUserTransactionLogsStmt, listFromUserTransactionLogs, fromUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TransactionRecord{}
	for rows.Next() {
		var i TransactionRecord
		if err := rows.Scan(
			&i.RowID,
			&i.FromUserID,
			&i.FromWalletID,
			&i.ToUserID,
			&i.ToWalletID,
			&i.Coin,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPaginatedFromUserTransactionLogs = `-- name: ListPaginatedFromUserTransactionLogs :many
SELECT row_id, from_user_id, from_wallet_id, to_user_id, to_wallet_id, coin, amount, created_at
FROM transaction_record
WHERE from_user_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3
`

type ListPaginatedFromUserTransactionLogsParams struct {
	FromUserID uuid.UUID `json:"fromUserID"`
	Limit      int32     `json:"limit"`
	Offset     int32     `json:"offset"`
}

func (q *Queries) ListPaginatedFromUserTransactionLogs(ctx context.Context, arg ListPaginatedFromUserTransactionLogsParams) ([]TransactionRecord, error) {
	rows, err := q.query(ctx, q.listPaginatedFromUserTransactionLogsStmt, listPaginatedFromUserTransactionLogs, arg.FromUserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TransactionRecord{}
	for rows.Next() {
		var i TransactionRecord
		if err := rows.Scan(
			&i.RowID,
			&i.FromUserID,
			&i.FromWalletID,
			&i.ToUserID,
			&i.ToWalletID,
			&i.Coin,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPaginatedToUserTransactionLogs = `-- name: ListPaginatedToUserTransactionLogs :many
SELECT row_id, from_user_id, from_wallet_id, to_user_id, to_wallet_id, coin, amount, created_at
FROM transaction_record
WHERE to_user_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3
`

type ListPaginatedToUserTransactionLogsParams struct {
	ToUserID uuid.UUID `json:"toUserID"`
	Limit    int32     `json:"limit"`
	Offset   int32     `json:"offset"`
}

func (q *Queries) ListPaginatedToUserTransactionLogs(ctx context.Context, arg ListPaginatedToUserTransactionLogsParams) ([]TransactionRecord, error) {
	rows, err := q.query(ctx, q.listPaginatedToUserTransactionLogsStmt, listPaginatedToUserTransactionLogs, arg.ToUserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TransactionRecord{}
	for rows.Next() {
		var i TransactionRecord
		if err := rows.Scan(
			&i.RowID,
			&i.FromUserID,
			&i.FromWalletID,
			&i.ToUserID,
			&i.ToWalletID,
			&i.Coin,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPaginatedTransactionLogs = `-- name: ListPaginatedTransactionLogs :many
SELECT row_id, from_user_id, from_wallet_id, to_user_id, to_wallet_id, coin, amount, created_at
FROM transaction_record
ORDER BY row_id
LIMIT $1 OFFSET $2
`

type ListPaginatedTransactionLogsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPaginatedTransactionLogs(ctx context.Context, arg ListPaginatedTransactionLogsParams) ([]TransactionRecord, error) {
	rows, err := q.query(ctx, q.listPaginatedTransactionLogsStmt, listPaginatedTransactionLogs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TransactionRecord{}
	for rows.Next() {
		var i TransactionRecord
		if err := rows.Scan(
			&i.RowID,
			&i.FromUserID,
			&i.FromWalletID,
			&i.ToUserID,
			&i.ToWalletID,
			&i.Coin,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listToUserTransactionLogs = `-- name: ListToUserTransactionLogs :many
SELECT row_id, from_user_id, from_wallet_id, to_user_id, to_wallet_id, coin, amount, created_at
FROM transaction_record
WHERE to_user_id = $1
ORDER BY row_id
`

func (q *Queries) ListToUserTransactionLogs(ctx context.Context, toUserID uuid.UUID) ([]TransactionRecord, error) {
	rows, err := q.query(ctx, q.listToUserTransactionLogsStmt, listToUserTransactionLogs, toUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TransactionRecord{}
	for rows.Next() {
		var i TransactionRecord
		if err := rows.Scan(
			&i.RowID,
			&i.FromUserID,
			&i.FromWalletID,
			&i.ToUserID,
			&i.ToWalletID,
			&i.Coin,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransactionLogs = `-- name: ListTransactionLogs :many
SELECT row_id, from_user_id, from_wallet_id, to_user_id, to_wallet_id, coin, amount, created_at
FROM transaction_record
ORDER BY row_id
`

func (q *Queries) ListTransactionLogs(ctx context.Context) ([]TransactionRecord, error) {
	rows, err := q.query(ctx, q.listTransactionLogsStmt, listTransactionLogs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TransactionRecord{}
	for rows.Next() {
		var i TransactionRecord
		if err := rows.Scan(
			&i.RowID,
			&i.FromUserID,
			&i.FromWalletID,
			&i.ToUserID,
			&i.ToWalletID,
			&i.Coin,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const logTransaction = `-- name: LogTransaction :one
INSERT INTO transaction_record
(from_user_id, from_wallet_id, to_user_id, to_wallet_id, coin, amount, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING row_id, from_user_id, from_wallet_id, to_user_id, to_wallet_id, coin, amount, created_at
`

type LogTransactionParams struct {
	FromUserID   uuid.UUID   `json:"fromUserID"`
	FromWalletID uuid.UUID   `json:"fromWalletID"`
	ToUserID     uuid.UUID   `json:"toUserID"`
	ToWalletID   uuid.UUID   `json:"toWalletID"`
	Coin         CryptoCoins `json:"coin"`
	Amount       int64       `json:"amount"`
	CreatedAt    time.Time   `json:"createdAt"`
}

func (q *Queries) LogTransaction(ctx context.Context, arg LogTransactionParams) (TransactionRecord, error) {
	row := q.queryRow(ctx, q.logTransactionStmt, logTransaction,
		arg.FromUserID,
		arg.FromWalletID,
		arg.ToUserID,
		arg.ToWalletID,
		arg.Coin,
		arg.Amount,
		arg.CreatedAt,
	)
	var i TransactionRecord
	err := row.Scan(
		&i.RowID,
		&i.FromUserID,
		&i.FromWalletID,
		&i.ToUserID,
		&i.ToWalletID,
		&i.Coin,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
