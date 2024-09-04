package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// store  provides all functions to execute db Queries and transactions
type Store struct {
	*Queries
	db *pgxpool.Pool
}

// NewStore create a new store
func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
