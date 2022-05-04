-- name: GetUserByID :one
SELECT *
FROM users
WHERE user_id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY row_id;

-- name: ListPaginatedUsers :many
SELECT *
FROM users
ORDER BY row_id
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (user_id, name, email, hashed_password, role, permissions, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET (name, email, hashed_password, role, permissions, updated_at) = ($2, $3, $4, $5, $6, $7)
WHERE user_id = $1
RETURNING *;

-- name: UpdateUserName :one
UPDATE users
SET (name, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING *;

-- name: UpdateUserEmail :one
UPDATE users
SET (email, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING *;

-- name: UpdateUserRole :one
UPDATE users
SET (role, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING *;

-- name: UpdateUserPermissions :one
UPDATE users
SET (permissions, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING *;

-- name: UpdateUserHashedPassword :one
UPDATE users
SET (hashed_password, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING *;

-- name: DeleteUserByID :exec
DELETE
FROM users
WHERE user_id = $1;

-- name: DeleteUserByEmail :exec
DELETE
FROM users
WHERE email = $1;
