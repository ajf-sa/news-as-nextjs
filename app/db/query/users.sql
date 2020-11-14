-- name: GetOneUSer :one
SELECT * FROM users WHERE username=$1 LIMIT 1 OFFSET 0 ;

-- name: AddNewUser :one
INSERT INTO users(username,password)VALUES($1,$2) RETURNING *;

-- name: ListUsers :many
SELECT * FROM users LIMIT $1 OFFSET $2;