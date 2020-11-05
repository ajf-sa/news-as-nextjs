-- name: AddNewPost :one
INSERT INTO posts(title,body)VALUES($1,$2) RETURNING *;

-- name: ListPosts :many
SELECT * FROM posts;

-- name: GetOnePost :one
SELECT * FROM posts WHERE id=$1 LIMIT 1 OFFSET 0 ;