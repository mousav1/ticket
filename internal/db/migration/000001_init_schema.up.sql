-- init_schema.up.sql

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    password_changed_at TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE terminals (
    id SERIAL PRIMARY KEY,
    city_id INT NOT NULL REFERENCES cities(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    UNIQUE(city_id, name)
);

CREATE TABLE routes (
    id SERIAL PRIMARY KEY,
    origin_terminal_id INT NOT NULL REFERENCES terminals(id) ON DELETE CASCADE,
    destination_terminal_id INT NOT NULL REFERENCES terminals(id) ON DELETE CASCADE,
    duration INTERVAL NOT NULL,
    distance INT NOT NULL CHECK (distance > 0),
    UNIQUE(origin_terminal_id, destination_terminal_id)
);

CREATE TABLE buses (
    id SERIAL PRIMARY KEY,
    route_id INT NOT NULL REFERENCES routes(id) ON DELETE CASCADE,
    departure_time TIMESTAMPTZ NOT NULL,
    arrival_time TIMESTAMPTZ NOT NULL,
    capacity INT NOT NULL CHECK (capacity > 0),
    price INT NOT NULL CHECK (price > 0),
    bus_type VARCHAR(255) NOT NULL,
    corporation VARCHAR(255),
    super_corporation VARCHAR(255),
    service_number VARCHAR(255),
    is_vip BOOLEAN DEFAULT FALSE
);

CREATE TABLE bus_seats (
    id SERIAL PRIMARY KEY,
    bus_id INT NOT NULL REFERENCES buses(id) ON DELETE CASCADE,
    seat_number INT NOT NULL CHECK (seat_number > 0),
    status VARCHAR(20) NOT NULL DEFAULT 'available' CHECK (status IN ('available', 'reserved', 'purchased', 'maintenance', 'broken'))
);

CREATE TABLE seat_reservations (
    id SERIAL PRIMARY KEY,
    bus_id INT NOT NULL REFERENCES buses(id) ON DELETE CASCADE,
    bus_seat_id INT NOT NULL REFERENCES bus_seats(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL CHECK (status IN ('reserved', 'purchased', 'canceled')),
    reserved_at TIMESTAMPTZ DEFAULT NOW(),
    purchased_at TIMESTAMPTZ -- Optional: to track when the reservation was turned into a purchase
);

CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    bus_id INT NOT NULL REFERENCES buses(id) ON DELETE CASCADE,
    seat_reservation_id INT NOT NULL REFERENCES seat_reservations(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL CHECK (status IN ('reserved', 'purchased', 'canceled')),
    reserved_at TIMESTAMPTZ DEFAULT NOW(),
    purchased_at TIMESTAMPTZ, -- Optional: to track when the reservation was turned into a purchase
    UNIQUE(user_id, bus_id, seat_reservation_id)
);

CREATE TABLE penalties (
    id SERIAL PRIMARY KEY,
    bus_id INT NOT NULL REFERENCES buses(id) ON DELETE CASCADE,
    actual_hours_before FLOAT CHECK (actual_hours_before >= 0),
    hours_before FLOAT CHECK (hours_before >= 0),
    percent INT NOT NULL CHECK (percent >= 0 AND percent <= 100),
    custom_text TEXT
);

CREATE TABLE sessions (
  id UUID PRIMARY KEY,
  username VARCHAR NOT NULL,
  refresh_token VARCHAR NOT NULL,
  user_agent VARCHAR NOT NULL,
  client_ip VARCHAR NOT NULL,
  is_blocked BOOLEAN NOT NULL DEFAULT false,
  expires_at TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  FOREIGN KEY (username) REFERENCES users (username) ON DELETE CASCADE
);

CREATE INDEX idx_routes_origin_destination ON routes (origin_terminal_id, destination_terminal_id);
CREATE INDEX idx_buses_departure_time ON buses (departure_time);
-- CREATE INDEX idx_bus_seats_status ON bus_seats (status);
CREATE INDEX idx_tickets_reserved_at ON tickets (reserved_at);
CREATE INDEX idx_sessions_username ON sessions (username);