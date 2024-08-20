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