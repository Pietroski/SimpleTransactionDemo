-- name: LogEntry :one
INSERT INTO entry_record
    (account_id, wallet_id, coin, amount, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListEntryLogs :many
SELECT *
FROM entry_record
ORDER BY row_id;

-- name: ListPaginatedEntryLogs :many
SELECT *
FROM entry_record
ORDER BY row_id
LIMIT $1 OFFSET $2;

-- name: ListEntryLogsByAccountID :many
SELECT *
FROM entry_record
WHERE account_id = $1
ORDER BY row_id;

-- name: ListPaginatedEntryLogsByAccountID :many
SELECT *
FROM entry_record
WHERE account_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3;

-- name: ListCoinEntryLogsByAccountID :many
SELECT *
FROM entry_record
WHERE account_id = $1
  AND coin = $2
ORDER BY row_id;

-- name: ListPaginatedCoinEntryLogsByAccountID :many
SELECT *
FROM entry_record
WHERE account_id = $1
  AND coin = $2
ORDER BY row_id
LIMIT $3 OFFSET $4;
