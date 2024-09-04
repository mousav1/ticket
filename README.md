# Ticket Booking System 🎫

This repository contains the final project of the Go (Golang) course available on our YouTube channel. The project demonstrates a complete Ticket Booking System built with Go, showcasing advanced features and best practices.

## 🚀 Project Overview

The Ticket Booking System is a web application built with Go, designed to manage ticket reservations. It includes features such as:

- **User Registration and Authentication:** Secure sign-up and login system for users.
- **Ticket Booking:** Users can book, view, and manage their tickets.
- **Concurrency:** Efficient handling of multiple requests using Go's goroutines and channels.
- **...**

## 🛠️ Technologies Used

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **Web Framework:** fiber
- **Authentication:** JWT (JSON Web Tokens)
- **Docker:** For containerization and deployment
- **golang-migrate:** Database migrations. CLI and Golang library.

## 📂 Project Structure

```plaintext
ticket/
├── cmd/             # Main application entry point
├── config/          # Configuration files
├── internal/        # Core application logic
│   ├── handlers/    # HTTP request handlers
│   ├── models/      # Database models
│   ├── services/    # Business logic services
├── migrations/      # Database migration files
├── scripts/         # Utility scripts
├── Dockerfile       # Docker configuration
└── README.md        # Project documentation
```

## 🛠️ Installation and Setup

### Clone the repository:
To get started with the project, follow these steps:

```bash
git clone https://github.com/mousav1/ticket.git
cd ticket

```
### Set up the environment:

Create a .env file based on .env.example and configure your database credentials and other settings.

### Run database migrations:

```bash
go run scripts/migrate.go
```

### Build and run the application:

```bash
migrate -path interna/db/migration -database "$(DB_URL)" up
```

### Using Docker (Optional):

```bash
docker-compose up --build
```

## 🚀 Usage

Access the application at http://localhost:8080.
Use the admin panel to create events and manage users.
Book tickets and explore the functionalities.
