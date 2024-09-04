-- buses.sql

-- name: CreateBus :one
INSERT INTO buses (route_id, departure_time, arrival_time, capacity, price, bus_type, corporation, super_corporation, service_number, is_vip)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, route_id, departure_time, arrival_time, capacity, price, bus_type, corporation, super_corporation, service_number, is_vip;

-- name: SearchBuses :many
SELECT b.id, b.route_id, b.departure_time, b.arrival_time, b.capacity, b.price, b.bus_type, b.corporation, b.super_corporation, b.service_number, b.is_vip
FROM buses b
JOIN routes r ON b.route_id = r.id
WHERE r.origin_terminal_id = $1 AND r.destination_terminal_id = $2 AND b.departure_time >= $3;

-- name: SearchBusesByCities :many
SELECT b.id, b.route_id, b.departure_time, b.arrival_time, b.capacity, b.price, b.bus_type, b.corporation, b.super_corporation, b.service_number, b.is_vip
FROM buses b
JOIN routes r ON b.route_id = r.id
JOIN terminals t_origin ON r.origin_terminal_id = t_origin.id
JOIN terminals t_destination ON r.destination_terminal_id = t_destination.id
WHERE t_origin.city_id = $1 AND t_destination.city_id = $2 AND b.departure_time >= $3;


-- name: GetBusByID :one
SELECT id, route_id, departure_time, arrival_time, capacity, price, bus_type, corporation, super_corporation, service_number, is_vip
FROM buses
WHERE id = $1;


-- name: CreateBusSeat :one
INSERT INTO bus_seats (bus_id, seat_number, status, passenger_national_code)
VALUES ($1, $2, $3, $4)
RETURNING id, bus_id, seat_number, status, passenger_national_code;

-- name: GetBusSeats :many
SELECT id, bus_id, seat_number, status, passenger_national_code
FROM bus_seats
WHERE bus_id = $1;

-- name: ListAvailableSeats :many
SELECT 
    bs.id AS seat_id,
    bs.seat_number,
    bs.status
FROM 
    bus_seats bs
    JOIN buses b ON bs.bus_id = b.id
WHERE 
    bs.bus_id = $1
    AND bs.status = 0 -- Assuming status = 0 means the seat is available
ORDER BY 
    bs.seat_number;

-- name: GetSeatByID :one
SELECT 
    s.id AS seat_id,
    s.bus_id,
    s.seat_number,
    s.status,
    s.passenger_national_code
FROM 
    bus_seats s
WHERE 
    s.id = $1
    AND s.bus_id = $2
LIMIT 1;

-- name: GetAvailableSeatsForBus :many
SELECT
    bs.id AS seat_id,
    bs.seat_number,
    bs.status
FROM
    bus_seats bs
    JOIN buses b ON bs.bus_id = b.id
WHERE
    b.route_id = $1
    AND bs.bus_id = $2
    AND bs.status = 0 
ORDER BY
    bs.seat_number;


-- name: UpdateSeatStatus :exec
UPDATE bus_seats
SET 
    status = $1,
    passenger_national_code = $2
WHERE 
    id = $3;


-- name: CheckBusRouteAssociation :one
SELECT 
    b.id AS bus_id,
    r.id AS route_id
FROM 
    buses b
JOIN 
    routes r ON b.route_id = r.id
WHERE 
    b.id = $1  -- BusID
    AND r.id = $2  -- RouteID
LIMIT 1;


-- CheckSeatAvailability :one
SELECT 
    s.id AS seat_id, 
    s.status 
FROM 
    bus_seats s
LEFT JOIN 
    tickets t ON s.id = t.seat_id AND t.status IN ('reserved', 'purchased')
WHERE 
    s.id = $1 
    AND s.bus_id = $2
    AND t.id IS NULL;  -- Ensures no conflicting reservation or purchase exists