-- name: AddNewPost :one
INSERT INTO posts(title,body)VALUES($1,$2) RETURNING *;