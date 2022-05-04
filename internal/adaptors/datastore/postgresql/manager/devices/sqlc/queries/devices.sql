-- name: CreateDevice :one
INSERT INTO devices (device_id, user_id, device_name, device_brand)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateDevice :one
UPDATE devices
SET (device_id, user_id, device_name, device_brand, updated_at) = ($1, $2, $3, $4, $5)
WHERE device_id = $1
RETURNING *;

-- name: UpdateUserDevice :one
UPDATE devices
SET (device_id, user_id, device_name, device_brand, updated_at) = ($1, $2, $3, $4, $5)
WHERE device_id = $1 AND user_id = $2
RETURNING *;

-- name: GetDevice :one
SELECT *
FROM devices
WHERE device_id = $1
LIMIT 1;

-- name: GetUserDevices :one
SELECT *
FROM devices
WHERE user_id = $1
ORDER BY row_id;

-- name: GetPaginatedUserDevices :one
SELECT *
FROM devices
WHERE user_id = $1
ORDER BY row_id
LIMIT $2 OFFSET $3;

-- name: ListDevices :many
SELECT *
FROM devices
ORDER BY row_id;

-- name: ListPaginatedDevices :many
SELECT *
FROM devices
ORDER BY row_id
LIMIT $1 OFFSET $2;

-- name: DeleteDevice :exec
DELETE
FROM devices
WHERE device_id = $1;

-- name: DeleteUserDevices :exec
DELETE
FROM devices
WHERE user_id = $1;
