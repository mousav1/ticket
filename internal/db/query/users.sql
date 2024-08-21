-- users.sql

-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserByID :one
SELECT id, username, hashed_password
FROM users
WHERE id = $1;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  full_name = COALESCE(sqlc.narg(full_name), full_name)
WHERE
  username = sqlc.arg(username)
RETURNING *;