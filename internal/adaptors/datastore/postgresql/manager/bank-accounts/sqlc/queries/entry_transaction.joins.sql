-- -- name: LogTransactionWithEntries :one
-- WITH transaction_record
--          (from_account_id, from_wallet_id, to_account_id, to_wallet_id, coin, amount, created_at)
--          AS (VALUES (sqlc.arg(from_account_id), sqlc.arg(from_wallet_id), sqlc.arg(to_account_id),
--                      sqlc.arg(to_wallet_id),
--                      sqlc.arg(coin), sqlc.arg(amount),
--                      sqlc.arg(created_at))),
--      ins1 AS (
--          INSERT INTO entry_record
--              (account_id, wallet_id, coin, amount, created_at)
--              SELECT (from_account_id, from_wallet_id, coin, amount, created_at)
--              FROM transaction_record),
--      ins2 AS (
--          INSERT INTO entry_record
--              (account_id, wallet_id, coin, amount, created_at)
--              SELECT (to_account_id, to_wallet_id, coin, amount, created_at))
-- SELECT * FROM transaction_record;

--
--
--
--

-- name: LogTransactionWithEntriesSimplified :one
WITH first_insert AS (
    INSERT INTO entry_record
        (account_id, wallet_id, coin, amount, created_at)
        VALUES (sqlc.arg(from_account_id), sqlc.arg(from_wallet_id), sqlc.arg(coin), sqlc.arg(amount),
                sqlc.arg(created_at))
        RETURNING *),
     second_insert AS (
         INSERT INTO entry_record
             (account_id, wallet_id, coin, amount, created_at)
             VALUES (sqlc.arg(to_account_id), sqlc.arg(to_wallet_id), sqlc.arg(coin), sqlc.arg(amount),
                     sqlc.arg(created_at))
             RETURNING *)
INSERT
INTO transaction_record
(from_account_id, from_wallet_id, to_account_id, to_wallet_id, coin, amount, created_at)
VALUES (sqlc.arg(from_account_id), sqlc.arg(from_wallet_id), sqlc.arg(to_account_id), sqlc.arg(to_wallet_id),
        sqlc.arg(coin), sqlc.arg(amount),
        sqlc.arg(created_at))
RETURNING *;
