-- name: GetUserSessionByID :one
SELECT *
FROM user_session
WHERE user_id = $1
LIMIT 1;

-- name: GetUserSessionByAccessToken :one
SELECT *
FROM user_session
WHERE access_token = $1
LIMIT 1;

-- name: CreateUserSession :one
INSERT INTO user_session (user_id, access_token, role, permissions, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateUserSessionByUserID :one
UPDATE user_session
SET (access_token, role, permissions, updated_at) = ($2, $3, $4, $5)
WHERE user_id = $1
RETURNING *;

-- name: UpdateUserSessionByUserAccessToken :one
UPDATE user_session
SET (role, permissions, updated_at) = ($2, $3, $4)
WHERE access_token = $1
RETURNING *;

-- name: DeleteUserSessionByID :exec
DELETE
FROM user_session
WHERE user_id = $1;

-- name: DeleteUserSessionByAccessToken :exec
DELETE
FROM user_session
WHERE access_token = $1;
