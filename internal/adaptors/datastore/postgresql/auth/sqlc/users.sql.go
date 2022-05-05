// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package sqlc_auth_store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (user_id, name, email, hashed_password, role, permissions, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
`

type CreateUserParams struct {
	UserID         uuid.UUID   `json:"userID"`
	Name           string      `json:"name"`
	Email          string      `json:"email"`
	HashedPassword string      `json:"hashedPassword"`
	Role           Roles       `json:"role"`
	Permissions    Permissions `json:"permissions"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.UserID,
		arg.Name,
		arg.Email,
		arg.HashedPassword,
		arg.Role,
		arg.Permissions,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Users
	err := row.Scan(
		&i.RowID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUserByEmail = `-- name: DeleteUserByEmail :exec
DELETE
FROM users
WHERE email = $1
`

func (q *Queries) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := q.exec(ctx, q.deleteUserByEmailStmt, deleteUserByEmail, email)
	return err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE
FROM users
WHERE user_id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, userID uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteUserByIDStmt, deleteUserByID, userID)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (Users, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i Users
	err := row.Scan(
		&i.RowID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
FROM users
WHERE user_id = $1
LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, userID uuid.UUID) (Users, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, userID)
	var i Users
	err := row.Scan(
		&i.RowID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPaginatedUsers = `-- name: ListPaginatedUsers :many
SELECT row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
FROM users
ORDER BY row_id
LIMIT $1 OFFSET $2
`

type ListPaginatedUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPaginatedUsers(ctx context.Context, arg ListPaginatedUsersParams) ([]Users, error) {
	rows, err := q.query(ctx, q.listPaginatedUsersStmt, listPaginatedUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Users{}
	for rows.Next() {
		var i Users
		if err := rows.Scan(
			&i.RowID,
			&i.UserID,
			&i.Name,
			&i.Email,
			&i.HashedPassword,
			&i.Role,
			&i.Permissions,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
FROM users
ORDER BY row_id
`

func (q *Queries) ListUsers(ctx context.Context) ([]Users, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Users{}
	for rows.Next() {
		var i Users
		if err := rows.Scan(
			&i.RowID,
			&i.UserID,
			&i.Name,
			&i.Email,
			&i.HashedPassword,
			&i.Role,
			&i.Permissions,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET (name, email, hashed_password, role, permissions, updated_at) = ($2, $3, $4, $5, $6, $7)
WHERE user_id = $1
RETURNING row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
`

type UpdateUserParams struct {
	UserID         uuid.UUID   `json:"userID"`
	Name           string      `json:"name"`
	Email          string      `json:"email"`
	HashedPassword string      `json:"hashedPassword"`
	Role           Roles       `json:"role"`
	Permissions    Permissions `json:"permissions"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error) {
	row := q.queryRow(ctx, q.updateUserStmt, updateUser,
		arg.UserID,
		arg.Name,
		arg.Email,
		arg.HashedPassword,
		arg.Role,
		arg.Permissions,
		arg.UpdatedAt,
	)
	var i Users
	err := row.Scan(
		&i.RowID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserEmail = `-- name: UpdateUserEmail :one
UPDATE users
SET (email, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
`

type UpdateUserEmailParams struct {
	UserID    uuid.UUID `json:"userID"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (Users, error) {
	row := q.queryRow(ctx, q.updateUserEmailStmt, updateUserEmail, arg.UserID, arg.Email, arg.UpdatedAt)
	var i Users
	err := row.Scan(
		&i.RowID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserHashedPassword = `-- name: UpdateUserHashedPassword :one
UPDATE users
SET (hashed_password, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
`

type UpdateUserHashedPasswordParams struct {
	UserID         uuid.UUID `json:"userID"`
	HashedPassword string    `json:"hashedPassword"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (q *Queries) UpdateUserHashedPassword(ctx context.Context, arg UpdateUserHashedPasswordParams) (Users, error) {
	row := q.queryRow(ctx, q.updateUserHashedPasswordStmt, updateUserHashedPassword, arg.UserID, arg.HashedPassword, arg.UpdatedAt)
	var i Users
	err := row.Scan(
		&i.RowID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserName = `-- name: UpdateUserName :one
UPDATE users
SET (name, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
`

type UpdateUserNameParams struct {
	UserID    uuid.UUID `json:"userID"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (Users, error) {
	row := q.queryRow(ctx, q.updateUserNameStmt, updateUserName, arg.UserID, arg.Name, arg.UpdatedAt)
	var i Users
	err := row.Scan(
		&i.RowID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserPermissions = `-- name: UpdateUserPermissions :one
UPDATE users
SET (permissions, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
`

type UpdateUserPermissionsParams struct {
	UserID      uuid.UUID   `json:"userID"`
	Permissions Permissions `json:"permissions"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

func (q *Queries) UpdateUserPermissions(ctx context.Context, arg UpdateUserPermissionsParams) (Users, error) {
	row := q.queryRow(ctx, q.updateUserPermissionsStmt, updateUserPermissions, arg.UserID, arg.Permissions, arg.UpdatedAt)
	var i Users
	err := row.Scan(
		&i.RowID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE users
SET (role, updated_at) = ($2, $3)
WHERE user_id = $1
RETURNING row_id, user_id, name, email, hashed_password, role, permissions, created_at, updated_at
`

type UpdateUserRoleParams struct {
	UserID    uuid.UUID `json:"userID"`
	Role      Roles     `json:"role"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (Users, error) {
	row := q.queryRow(ctx, q.updateUserRoleStmt, updateUserRole, arg.UserID, arg.Role, arg.UpdatedAt)
	var i Users
	err := row.Scan(
		&i.RowID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
