-- UpdateAccountWalletBalance - deltas the user's balance within the database so it avoids potential deadlocks.
-- name: UpdateAccountWalletBalance :one
UPDATE wallet
SET balance    = balance + sqlc.arg(amount),
    updated_at = sqlc.arg(updated_at)
WHERE account_id = sqlc.arg(account_id)
  AND wallet_id = sqlc.arg(wallet_id)
  AND coin = sqlc.arg(coin)
RETURNING *;

-- GetTxWallet - gets the transacting user's wallet.
-- name: GetTxWallet :one
SELECT *
FROM wallet
WHERE account_id = $1
  AND coin = $2;

-- GetWalletsByAccountID - gets the user's wallet.
-- name: GetWalletsByAccountID :many
SELECT *
FROM wallet
WHERE account_id = $1
ORDER BY row_id;

-- GetPaginatedWalletsByAccountID - paginates the return of the user's wallet.
-- name: GetPaginatedWalletsByAccountID :many
SELECT *
FROM wallet
WHERE account_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3;

-- GetAccountWallets - deprecated method
-- name: GetAccountWallets :many
SELECT w.*
FROM account a
         INNER JOIN wallet w
                    ON a.account_id = w.account_id
WHERE a.account_id = $1
ORDER BY w.row_id;

-- GetPaginatedAccountWallets - deprecated method
-- name: GetPaginatedAccountWallets :many
SELECT w.*
FROM account a
         INNER JOIN wallet w
                    ON a.account_id = w.account_id
WHERE a.account_id = $1
ORDER BY w.row_id
LIMIT $2 OFFSET $3;
