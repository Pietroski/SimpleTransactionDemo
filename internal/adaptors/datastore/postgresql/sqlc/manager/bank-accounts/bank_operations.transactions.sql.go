// Code generated by sqlc. DO NOT EDIT.
// source: bank_operations.transactions.sql

package bank_account_store

import (
	"context"

	"github.com/google/uuid"
)

const getPaginatedUserWallets = `-- name: GetPaginatedUserWallets :many
SELECT w.row_id, w.wallet_id, w.user_id, w.coin, w.balance, w.created_at, w.updated_at
FROM account a
         INNER JOIN wallet w
                    ON a.user_id = w.user_id
WHERE a.user_id = $1
ORDER BY w.row_id
LIMIT $2 OFFSET $3
`

type GetPaginatedUserWalletsParams struct {
	UserID uuid.UUID `json:"userID"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

// GetPaginatedUserWallets - deprecated method
func (q *Queries) GetPaginatedUserWallets(ctx context.Context, arg GetPaginatedUserWalletsParams) ([]Wallet, error) {
	rows, err := q.query(ctx, q.getPaginatedUserWalletsStmt, getPaginatedUserWallets, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Wallet{}
	for rows.Next() {
		var i Wallet
		if err := rows.Scan(
			&i.RowID,
			&i.WalletID,
			&i.UserID,
			&i.Coin,
			&i.Balance,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getPaginatedWalletsByUserID = `-- name: GetPaginatedWalletsByUserID :many
SELECT row_id, wallet_id, user_id, coin, balance, created_at, updated_at
FROM wallet
WHERE user_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3
`

type GetPaginatedWalletsByUserIDParams struct {
	UserID uuid.UUID `json:"userID"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

// GetPaginatedWalletsByUserID - paginates the return of the user's wallet.
func (q *Queries) GetPaginatedWalletsByUserID(ctx context.Context, arg GetPaginatedWalletsByUserIDParams) ([]Wallet, error) {
	rows, err := q.query(ctx, q.getPaginatedWalletsByUserIDStmt, getPaginatedWalletsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Wallet{}
	for rows.Next() {
		var i Wallet
		if err := rows.Scan(
			&i.RowID,
			&i.WalletID,
			&i.UserID,
			&i.Coin,
			&i.Balance,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getUserWallets = `-- name: GetUserWallets :many
SELECT w.row_id, w.wallet_id, w.user_id, w.coin, w.balance, w.created_at, w.updated_at
FROM account a
         INNER JOIN wallet w
                    ON a.user_id = w.user_id
WHERE a.user_id = $1
ORDER BY w.row_id
`

// GetUserWallets - deprecated method
func (q *Queries) GetUserWallets(ctx context.Context, userID uuid.UUID) ([]Wallet, error) {
	rows, err := q.query(ctx, q.getUserWalletsStmt, getUserWallets, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Wallet{}
	for rows.Next() {
		var i Wallet
		if err := rows.Scan(
			&i.RowID,
			&i.WalletID,
			&i.UserID,
			&i.Coin,
			&i.Balance,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getWalletsByUserID = `-- name: GetWalletsByUserID :many
SELECT row_id, wallet_id, user_id, coin, balance, created_at, updated_at
FROM wallet
WHERE user_id = $1
ORDER BY row_id
`

// GetWalletsByUserID - gets the user's wallet.
func (q *Queries) GetWalletsByUserID(ctx context.Context, userID uuid.UUID) ([]Wallet, error) {
	rows, err := q.query(ctx, q.getWalletsByUserIDStmt, getWalletsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Wallet{}
	for rows.Next() {
		var i Wallet
		if err := rows.Scan(
			&i.RowID,
			&i.WalletID,
			&i.UserID,
			&i.Coin,
			&i.Balance,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateUserWalletBalance = `-- name: UpdateUserWalletBalance :one
UPDATE wallet
SET (balance) = balance + $1
WHERE user_id = $2
  AND coin = $3
RETURNING row_id, wallet_id, user_id, coin, balance, created_at, updated_at
`

type UpdateUserWalletBalanceParams struct {
	Amount int64       `json:"amount"`
	UserID uuid.UUID   `json:"userID"`
	Coin   CryptoCoins `json:"coin"`
}

// UpdateUserWalletBalance - deltas the user's balance within the database so it avoids potential deadlocks.
func (q *Queries) UpdateUserWalletBalance(ctx context.Context, arg UpdateUserWalletBalanceParams) (Wallet, error) {
	row := q.queryRow(ctx, q.updateUserWalletBalanceStmt, updateUserWalletBalance, arg.Amount, arg.UserID, arg.Coin)
	var i Wallet
	err := row.Scan(
		&i.RowID,
		&i.WalletID,
		&i.UserID,
		&i.Coin,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
