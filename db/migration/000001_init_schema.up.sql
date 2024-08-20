-- init_schema.up.sql

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE terminals (
    id SERIAL PRIMARY KEY,
    city_id INT REFERENCES cities(id),
    name VARCHAR(255) NOT NULL,
    UNIQUE(city_id, name)
);

CREATE TABLE routes (
    id SERIAL PRIMARY KEY,
    origin_terminal_id INT REFERENCES terminals(id),
    destination_terminal_id INT REFERENCES terminals(id),
    duration INTERVAL NOT NULL,
    distance INT NOT NULL,
    UNIQUE(origin_terminal_id, destination_terminal_id)
);

CREATE TABLE buses (
    id SERIAL PRIMARY KEY,
    route_id INT REFERENCES routes(id),
    departure_time TIMESTAMPTZ NOT NULL,
    arrival_time TIMESTAMPTZ NOT NULL,
    capacity INT NOT NULL,
    price INT NOT NULL,
    bus_type VARCHAR(255) NOT NULL,
    corporation VARCHAR(255),
    super_corporation VARCHAR(255),
    service_number VARCHAR(255),
    is_vip BOOLEAN DEFAULT FALSE
);

CREATE TABLE bus_seats (
    id SERIAL PRIMARY KEY,
    bus_id INT REFERENCES buses(id),
    seat_number INT NOT NULL,
    status INT NOT NULL,
    passenger_national_code VARCHAR(255)
);

CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    bus_id INT REFERENCES buses(id),
    seat_id INT REFERENCES bus_seats(id),
    reserved_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(user_id, bus_id, seat_id)
);

CREATE TABLE penalties (
    id SERIAL PRIMARY KEY,
    bus_id INT REFERENCES buses(id),
    actual_hours_before FLOAT,
    hours_before FLOAT,
    percent INT NOT NULL,
    custom_text TEXT
);