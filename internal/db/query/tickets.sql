-- name: GetTicketByID :one
SELECT id, user_id, bus_id, reserved_at
FROM tickets
WHERE id = $1;

-- name: DeleteTicket :exec
DELETE FROM tickets
WHERE id = $1;

-- name: GetReservedTicketsCount :one
SELECT COUNT(*)
FROM tickets
WHERE bus_id = $1;

-- name: ReserveTicket :exec
INSERT INTO tickets (user_id, bus_id, seat_id)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, bus_id, seat_id) DO NOTHING;

-- name: GetUserTickets :many
SELECT t.id, b.route_id, b.departure_time, b.arrival_time, b.capacity, b.price, b.bus_type, b.corporation, b.super_corporation, b.service_number, b.is_vip
FROM tickets t
JOIN buses b ON t.bus_id = b.id
WHERE t.user_id = $1;

-- name: CreatePenalty :one
INSERT INTO penalties (bus_id, actual_hours_before, hours_before, percent, custom_text)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, bus_id, actual_hours_before, hours_before, percent, custom_text;

-- name: GetBusPenalties :many
SELECT id, bus_id, actual_hours_before, hours_before, percent, custom_text
FROM penalties
WHERE bus_id = $1;