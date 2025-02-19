-- name: CreateTodo :one
INSERT INTO todos (title, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY created_at DESC;

-- name: UpdateTodo :one
UPDATE todos
SET
    title = COALESCE($2, title),
    description = COALESCE($3, description),
    completed = COALESCE($4, completed),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;
