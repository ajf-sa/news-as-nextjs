package db

import (
	"context"
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

func (q *Queries) SetPost(arg AddNewPostParams) (Post, error) {

	return q.AddNewPost(context.Background(), arg)
}
