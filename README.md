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
- **sqlc:** Generates Go code from SQL queries for type-safe database interactions

## 📂 Project Structure

```plaintext

ticket/
├── internal/                # Core application logic
│   ├── api/                 # HTTP request handling
│   │   ├── handlers/        # Controllers and request handlers
│   │   ├── middleware/      # Middleware for request processing
│   │   └── server.go        # Server initialization and configuration
│   ├── db/                  # Database connections, migrations, and models
│   ├── routes/              # Route definitions and management
│   ├── token/               # JWT token generation, validation, and authentication
│   ├── util/                # Utility functions, helpers, and shared components
├── Dockerfile               # Docker configuration for containerizing the application
├── docker-compose.yml       # Docker Compose configuration for orchestrating services
├── README.md                # Project documentation, setup instructions, and usage
├── main.go                  # Main entry point of the application
└── app.env                  # Environment variables configuration file


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
migrate -path interna/db/migration -database "postgres://username:password@localhost:5432/database_name?sslmode=disable" up
```

### Build and run the application:

```bash
go run main.go
```

### Using Docker (Optional):

```bash
docker-compose up --build
```

## 🚀 Usage

Access the application at http://localhost:8080.
Use the admin panel to create events and manage users.
Book tickets and explore the functionalities.
