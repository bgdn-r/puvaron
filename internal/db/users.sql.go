// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users
(id, created_at, updated_at, deleted_at, email, phone, first_name, last_name, username, password)
VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, created_at, updated_at, deleted_at, email, phone, first_name, last_name, username, password
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Email     string
	Phone     sql.NullString
	FirstName string
	LastName  string
	Username  string
	Password  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Email,
		arg.Phone,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Phone,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Password,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE users
SET password = $2
WHERE id = $1
RETURNING id, created_at, updated_at, deleted_at, email, phone, first_name, last_name, username, password
`

type UpdatePasswordParams struct {
	ID       uuid.UUID
	Password string
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updatePassword, arg.ID, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Phone,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Password,
	)
	return i, err
}
