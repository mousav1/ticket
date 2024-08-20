-- routes.sql

-- name: CreateRoute :one
INSERT INTO routes (origin_terminal_id, destination_terminal_id, duration, distance)
VALUES ($1, $2, $3, $4)
RETURNING id, origin_terminal_id, destination_terminal_id, duration, distance;

-- پیدا کردن تمامی مسیرها
-- name: GetAllRoutes :many
SELECT id, origin_terminal_id, destination_terminal_id, duration, distance
FROM routes;

-- پیدا کردن یک مسیر بر اساس مبدا و مقصد
-- name: GetRouteByTerminals :one
SELECT id, origin_terminal_id, destination_terminal_id, duration, distance
FROM routes
WHERE origin_terminal_id = $1 AND destination_terminal_id = $2;
