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
INSERT INTO bus_seats (bus_id, seat_number, status)
VALUES ($1, $2, 'available')
RETURNING id, bus_id, seat_number, status;

-- name: GetBusSeats :many
SELECT id, bus_id, seat_number, status
FROM bus_seats
WHERE bus_id = $1;

-- name: GetSeatByID :one
SELECT 
    bs.id AS seat_id,                
    bs.bus_id,                  
    bs.seat_number, 
    bs.status AS seat_status,              
    sr.status AS reservation_status,
    sr.user_id
FROM 
    bus_seats bs
LEFT JOIN 
    seat_reservations sr ON bs.id = sr.bus_seat_id
WHERE 
    bs.id = $1
    AND bs.bus_id = $2 
LIMIT 1;

-- name: GetAvailableSeatsForBus :many
SELECT
    bs.id AS seat_id,
    bs.seat_number,
    bs.status
FROM
    bus_seats bs
JOIN 
    buses b ON bs.bus_id = b.id
WHERE
    b.route_id = $1
    AND bs.bus_id = $2
    AND bs.status = 'available' -- Only select seats that are available
ORDER BY
    bs.seat_number;



-- name: UpdateSeatReservationStatus :exec
UPDATE seat_reservations
SET 
    status = $1
WHERE 
    bus_seat_id = $2
    AND user_id = $3;


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

-- name: CheckSeatAvailability :one
SELECT 
    s.id AS seat_id, 
    s.status 
FROM 
    bus_seats s
LEFT JOIN 
    seat_reservations sr ON s.id = sr.bus_seat_id AND sr.status IN ('reserved', 'purchased')
WHERE 
    s.id = $1 
    AND s.bus_id = $2
    AND s.status = 'available' -- Ensures seat is available
    AND sr.id IS NULL;  -- Ensures no conflicting reservation or purchase exists


-- name: UpdateSeatStatusAfterTrip :exec
UPDATE bus_seats
SET status = 'available'
WHERE bus_id = $1
  AND status = 'purchased';
