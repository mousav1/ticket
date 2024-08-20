-- cities.sql

-- name: GetCityByID :one
SELECT id, name
FROM cities
WHERE id = $1;

-- name: CreateCity :one
INSERT INTO cities (name)
VALUES ($1)
RETURNING id, name;

-- name: GetAllCities :many
SELECT id, name
FROM cities;
