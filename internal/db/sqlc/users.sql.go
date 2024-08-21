// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one

INSERT INTO users (
  username,
  hashed_password,
  full_name
) VALUES (
  $1, $2, $3
) RETURNING id, username, hashed_password, full_name, password_changed_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
}

// users.sql
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.HashedPassword, arg.FullName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, hashed_password, full_name, password_changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, hashed_password
FROM users
WHERE id = $1
`

type GetUserByIDRow struct {
	ID             int32  `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) GetUserByID(ctx context.Context, id int32) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(&i.ID, &i.Username, &i.HashedPassword)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE($1, hashed_password),
  password_changed_at = COALESCE($2, password_changed_at),
  full_name = COALESCE($3, full_name)
WHERE
  username = $4
RETURNING id, username, hashed_password, full_name, password_changed_at, created_at
`

type UpdateUserParams struct {
	HashedPassword    pgtype.Text        `json:"hashed_password"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	FullName          pgtype.Text        `json:"full_name"`
	Username          string             `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.FullName,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}