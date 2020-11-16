-- name: AddNewPost :one
INSERT INTO posts(title,body,thumbnail,image)VALUES($1,$2,$3,$4) RETURNING *;

-- name: ListPosts :many
SELECT * FROM posts;

-- name: GetOnePost :one
SELECT * FROM posts WHERE id=$1 LIMIT 1 OFFSET 0 ;