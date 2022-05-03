-- name: GetUserRecoveryByID :one
SELECT *
FROM recovering
WHERE user_id = $1
LIMIT 1;

-- name: GetUserRecoveryByRecoveryLink :one
SELECT *
FROM recovering
WHERE recovery_link = $1
LIMIT 1;

-- name: CreateUserRecovery :one
INSERT INTO recovering (user_id, recovery_link, expiry_date, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUserRecoveryByUserID :one
UPDATE recovering
SET (recovery_link, expiry_date, updated_at) = ($2, $3, $4)
WHERE user_id = $1
RETURNING *;

-- name: UpdateUserRecoveryByRecoveryLink :one
UPDATE recovering
SET (expiry_date, updated_at) = ($2, $3)
WHERE recovery_link = $1
RETURNING *;

-- name: DeleteUserRecoveryByID :exec
DELETE
FROM recovering
WHERE user_id = $1;

-- name: DeleteUserRecoveryByRecoveryLink :exec
DELETE
FROM recovering
WHERE recovery_link = $1;
