-- terminals.sql

-- name: CreateTerminal :one
INSERT INTO terminals (city_id, name)
VALUES ($1, $2)
RETURNING id, city_id, name;

-- name: GetTerminalByID :one
SELECT id, city_id, name
FROM terminals
WHERE id = $1;

-- name: GetTerminalsByCity :many
SELECT id, city_id, name
FROM terminals
WHERE city_id = $1;

-- name: ListTerminals :many
SELECT id, name, city_id
FROM terminals
ORDER BY name;
