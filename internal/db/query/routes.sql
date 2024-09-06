-- routes.sql

-- name: CreateRoute :one
INSERT INTO routes (origin_terminal_id, destination_terminal_id, duration, distance)
VALUES ($1, $2, $3, $4)
RETURNING id, origin_terminal_id, destination_terminal_id, duration, distance;

-- name: GetAllRoutes :many
SELECT id, origin_terminal_id, destination_terminal_id, duration, distance
FROM routes;

-- name: GetRouteByID :one
SELECT id, origin_terminal_id, destination_terminal_id, duration, distance
FROM routes
WHERE id = $1;

-- name: ListRoutes :many
SELECT 
    r.id AS route_id,
    r.origin_terminal_id,
    r.destination_terminal_id,
    t1.name AS origin_terminal_name,
    t2.name AS destination_terminal_name,
    b.id AS bus_id,
    b.departure_time,
    b.arrival_time,
    b.capacity,
    b.price,
    b.bus_type,
    b.corporation,
    b.super_corporation,
    b.service_number,
    b.is_vip,
    -- Calculate available seats by counting seats in 'available' status
    COUNT(bs.id) FILTER (
        WHERE bs.status = 'available'
    ) AS available_seats
FROM 
    routes r
    JOIN terminals t1 ON r.origin_terminal_id = t1.id
    JOIN terminals t2 ON r.destination_terminal_id = t2.id
    JOIN buses b ON b.route_id = r.id
    LEFT JOIN bus_seats bs ON bs.bus_id = b.id
WHERE 
    r.origin_terminal_id = $1
    AND r.destination_terminal_id = $2
    AND b.departure_time::date = $3
GROUP BY 
    r.id, t1.name, t2.name, b.id;
