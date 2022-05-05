// Code generated by sqlc. DO NOT EDIT.

package sqlc_auth_store

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	CreateUserRecovery(ctx context.Context, arg CreateUserRecoveryParams) (Recovering, error)
	CreateUserSession(ctx context.Context, arg CreateUserSessionParams) (UserSession, error)
	DeleteUserByEmail(ctx context.Context, email string) error
	DeleteUserByID(ctx context.Context, userID uuid.UUID) error
	DeleteUserRecoveryByID(ctx context.Context, userID uuid.UUID) error
	DeleteUserRecoveryByRecoveryLink(ctx context.Context, recoveryLink string) error
	DeleteUserSessionByAccessToken(ctx context.Context, accessToken string) error
	DeleteUserSessionByID(ctx context.Context, userID uuid.UUID) error
	GetUserByEmail(ctx context.Context, email string) (Users, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (Users, error)
	GetUserRecoveryByID(ctx context.Context, userID uuid.UUID) (Recovering, error)
	GetUserRecoveryByRecoveryLink(ctx context.Context, recoveryLink string) (Recovering, error)
	GetUserSessionByAccessToken(ctx context.Context, accessToken string) (UserSession, error)
	GetUserSessionByID(ctx context.Context, userID uuid.UUID) (UserSession, error)
	ListPaginatedUsers(ctx context.Context, arg ListPaginatedUsersParams) ([]Users, error)
	ListUsers(ctx context.Context) ([]Users, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (Users, error)
	UpdateUserHashedPassword(ctx context.Context, arg UpdateUserHashedPasswordParams) (Users, error)
	UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (Users, error)
	UpdateUserPermissions(ctx context.Context, arg UpdateUserPermissionsParams) (Users, error)
	UpdateUserRecoveryByRecoveryLink(ctx context.Context, arg UpdateUserRecoveryByRecoveryLinkParams) (Recovering, error)
	UpdateUserRecoveryByUserID(ctx context.Context, arg UpdateUserRecoveryByUserIDParams) (Recovering, error)
	UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (Users, error)
	UpdateUserSessionByUserAccessToken(ctx context.Context, arg UpdateUserSessionByUserAccessTokenParams) (UserSession, error)
	UpdateUserSessionByUserID(ctx context.Context, arg UpdateUserSessionByUserIDParams) (UserSession, error)
}

var _ Querier = (*Queries)(nil)
