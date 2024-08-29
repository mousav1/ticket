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

-- name: GetUserByUsername :one
SELECT id, username, hashed_password
FROM users
WHERE username = $1;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  full_name = $1
WHERE
  username = $2
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
SET hashed_password = $1,
    password_changed_at = now()
WHERE username = $2
RETURNING *;
