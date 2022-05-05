-- name: LogTransaction :one
INSERT INTO transaction_record
(from_account_id, from_wallet_id, to_account_id, to_wallet_id, coin, amount, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: ListTransactionLogs :many
SELECT *
FROM transaction_record
ORDER BY row_id;

-- name: ListPaginatedTransactionLogs :many
SELECT *
FROM transaction_record
ORDER BY row_id
LIMIT $1 OFFSET $2;

-- name: ListFromAccountTransactionLogs :many
SELECT *
FROM transaction_record
WHERE from_account_id = $1
ORDER BY row_id;

-- name: ListPaginatedFromAccountTransactionLogs :many
SELECT *
FROM transaction_record
WHERE from_account_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3;

-- name: ListToAccountTransactionLogs :many
SELECT *
FROM transaction_record
WHERE to_account_id = $1
ORDER BY row_id;

-- name: ListPaginatedToAccountTransactionLogs :many
SELECT *
FROM transaction_record
WHERE to_account_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3;
