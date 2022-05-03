-- UpdateUserWalletBalance - deltas the user's balance within the database so it avoids potential deadlocks.
-- name: UpdateUserWalletBalance :one
UPDATE wallet
SET (balance) = balance + sqlc.arg(amount)
WHERE user_id = sqlc.arg(user_id)
  AND coin = sqlc.arg(coin)
RETURNING *;

-- GetWalletsByUserID - gets the user's wallet.
-- name: GetWalletsByUserID :many
SELECT *
FROM wallet
WHERE user_id = $1
ORDER BY row_id;

-- GetPaginatedWalletsByUserID - paginates the return of the user's wallet.
-- name: GetPaginatedWalletsByUserID :many
SELECT *
FROM wallet
WHERE user_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3;

-- GetUserWallets - deprecated method
-- name: GetUserWallets :many
SELECT w.*
FROM account a
         INNER JOIN wallet w
                    ON a.user_id = w.user_id
WHERE a.user_id = $1
ORDER BY w.row_id;

-- GetPaginatedUserWallets - deprecated method
-- name: GetPaginatedUserWallets :many
SELECT w.*
FROM account a
         INNER JOIN wallet w
                    ON a.user_id = w.user_id
WHERE a.user_id = $1
ORDER BY w.row_id
LIMIT $2 OFFSET $3;
