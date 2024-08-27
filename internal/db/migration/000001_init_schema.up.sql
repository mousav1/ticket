-- init_schema.up.sql

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    password_changed_at timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
    created_at timestamptz NOT NULL DEFAULT (now())
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

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");