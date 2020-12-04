// Code generated by sqlc. DO NOT EDIT.
// source: posts.sql

package db

import (
	"context"
)

const addNewPost = `-- name: AddNewPost :one
INSERT INTO posts(title,body)VALUES($1,$2) RETURNING id, title, descr, body, thumbnail, image, create_at, user_id, page_id
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
		&i.PageID,
	)
	return i, err
}

const getOnePost = `-- name: GetOnePost :one
SELECT id, title, descr, body, thumbnail, image, create_at, user_id, page_id FROM posts WHERE id=$1 LIMIT 1 OFFSET 0
`

func (q *Queries) GetOnePost(ctx context.Context, id int32) (Post, error) {
	row := q.queryRow(ctx, q.getOnePostStmt, getOnePost, id)
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
		&i.PageID,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id, title, descr, body, thumbnail, image, create_at, user_id, page_id FROM posts
`

func (q *Queries) ListPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.query(ctx, q.listPostsStmt, listPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Descr,
			&i.Body,
			&i.Thumbnail,
			&i.Image,
			&i.CreateAt,
			&i.UserID,
			&i.PageID,
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
