-- name: LogTransaction :one
INSERT INTO transaction_record
(from_user_id, from_wallet_id, to_user_id, to_wallet_id, coin, amount, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
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

-- name: ListFromUserTransactionLogs :many
SELECT *
FROM transaction_record
WHERE from_user_id = $1
ORDER BY row_id;

-- name: ListPaginatedFromUserTransactionLogs :many
SELECT *
FROM transaction_record
WHERE from_user_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3;

-- name: ListToUserTransactionLogs :many
SELECT *
FROM transaction_record
WHERE to_user_id = $1
ORDER BY row_id;

-- name: ListPaginatedToUserTransactionLogs :many
SELECT *
FROM transaction_record
WHERE to_user_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3;
