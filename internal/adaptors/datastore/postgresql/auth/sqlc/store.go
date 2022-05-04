package sqlc_auth_store

import (
	"database/sql"
)

type Store interface {
	Querier
}

type authStore struct {
	*Queries
	db *sql.DB
}

// NewStore returns a new Store interface
func NewStore(db *sql.DB) Store {
	store := &authStore{
		Queries: New(db),
		db:      db,
	}

	return store
}
