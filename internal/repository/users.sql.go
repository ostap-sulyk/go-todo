// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package repository

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password, created_at, updated_at
`

type CreateUserParams struct {
	Email    string
	Password string
}

// CREATE
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

// DELETE
func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const deleteUserByEmail = `-- name: DeleteUserByEmail :exec
DELETE FROM users WHERE email = $1
`

func (q *Queries) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, deleteUserByEmail, email)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, password, created_at, updated_at FROM users WHERE id = $1
`

// READ
func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserEmail = `-- name: UpdateUserEmail :exec
UPDATE users SET email = $2 WHERE id = $1
`

type UpdateUserEmailParams struct {
	ID    int64
	Email string
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error {
	_, err := q.db.Exec(ctx, updateUserEmail, arg.ID, arg.Email)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users SET password = $2 WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID       int64
	Password string
}

// UPDATE
func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.ID, arg.Password)
	return err
}
