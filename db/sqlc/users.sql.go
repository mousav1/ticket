// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one

INSERT INTO users (username, password_hash)
VALUES ($1, $2)
RETURNING id, username, password_hash
`

type CreateUserParams struct {
	Username     string
	PasswordHash string
}

// users.sql
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.PasswordHash)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.PasswordHash)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, password_hash
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.PasswordHash)
	return i, err
}