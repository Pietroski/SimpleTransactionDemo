package sqlc_device_store

import (
	"database/sql"
)

// Store provides all functions to execute db queries
type Store interface {
	Querier
}

type deviceStore struct {
	*Queries
	db *sql.DB
}

// NewStore instantiates a devices store object returning the store interface.
func NewStore(db *sql.DB) Store {
	store := &deviceStore{
		Queries: New(db),
		db:      db,
	}

	return store
}
