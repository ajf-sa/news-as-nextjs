// Code generated by sqlc. DO NOT EDIT.
// source: posts.sql

package db

import (
	"context"
)

const addNewPost = `-- name: AddNewPost :one
INSERT INTO posts(title,body)VALUES($1,$2) RETURNING id, title, descr, body, thumbnail, image, create_at, user_id, cate_id
`

type AddNewPostParams struct {
	Title string
	Body  string
}

func (q *Queries) AddNewPost(ctx context.Context, arg AddNewPostParams) (Post, error) {
	row := q.queryRow(ctx, q.addNewPostStmt, addNewPost, arg.Title, arg.Body)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Descr,
		&i.Body,
		&i.Thumbnail,
		&i.Image,
		&i.CreateAt,
		&i.UserID,
		&i.CateID,
	)
	return i, err
}