package db

import (
	"database/sql"
)

// Entiry struct
type Entiry struct {
	*Queries
	DB *sql.DB
}

// NewEntiry create new Entiry
func NewEntiry(db *sql.DB) *Entiry {
	return &Entiry{
		Queries: New(db),
		DB:      db,
	}
}
