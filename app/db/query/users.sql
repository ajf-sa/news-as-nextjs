-- name: GetOneUSer :one
SELECT * FROM users WHERE username=$1 and is_active=true LIMIT 1 OFFSET 0   ;

-- name: AddNewUser :one
INSERT INTO users(username,password,email)VALUES($1,$2,$3) RETURNING *;

-- name: ListUsers :many
SELECT * FROM users LIMIT $1 OFFSET $2;

-- name: EnableUser :exec
UPDATE users SET is_active=true;

-- name: DisableUser :exec
UPDATE users SET is_active=false;