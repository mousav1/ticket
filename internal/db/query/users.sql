-- users.sql

-- name: CreateUser :one
INSERT INTO users (username, password_hash)
VALUES ($1, $2)
RETURNING id, username, password_hash;

-- name: GetUserByID :one
SELECT id, username, password_hash
FROM users
WHERE id = $1;
