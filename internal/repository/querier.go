// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	// CREATE
	CreateTodo(ctx context.Context, arg CreateTodoParams) error
	// CREATE
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	// DELETE
	DeleteTodo(ctx context.Context, id int64) error
	DeleteTodosByUserId(ctx context.Context, userID pgtype.Int8) error
	// DELETE
	DeleteUser(ctx context.Context, id int64) error
	DeleteUserByEmail(ctx context.Context, email string) error
	// READ
	GetTodoById(ctx context.Context, id int64) (Todo, error)
	GetTodosByUserId(ctx context.Context, userID pgtype.Int8) ([]Todo, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	// READ
	GetUserById(ctx context.Context, id int64) (User, error)
	// UPDATE
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) error
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error
	// UPDATE
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
