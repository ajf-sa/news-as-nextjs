// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Post struct {
	ID        int32
	Title     string
	Descr     string
	Body      string
	Thumbnail string
	Image     string
	CreateAt  time.Time
	UserID    int32
	CateID    int32
}