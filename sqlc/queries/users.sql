-- CREATE
-- name: CreateUser :one
INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *;

-- READ
-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- UPDATE
-- name: UpdateUserPassword :exec
UPDATE users SET password = $2 WHERE id = $1;

-- name: UpdateUserEmail :exec
UPDATE users SET email = $2 WHERE id = $1;

-- DELETE
-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: DeleteUserByEmail :exec
DELETE FROM users WHERE email = $1;