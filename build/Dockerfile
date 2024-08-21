# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

# Stage 2: Create a minimal Docker image with the Go binary
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

CMD ["./main"]
