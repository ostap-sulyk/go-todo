-- CREATE
-- name: CreateTodo :exec
INSERT INTO todos (title, description, user_id) VALUES ($1, $2, $3);

-- READ
-- name: GetTodoById :one
SELECT * FROM todos WHERE id = $1;

-- name: GetTodosByUserId :many
SELECT * FROM todos WHERE user_id = $1;

-- UPDATE
-- name: UpdateTodo :exec
UPDATE todos SET title = $2, description = $3 WHERE id = $1;

-- DELETE
-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;

-- name: DeleteTodosByUserId :exec
DELETE FROM todos WHERE user_id = $1;