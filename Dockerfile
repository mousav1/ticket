# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -o main main.go

# Stage 2: Create a minimal Docker image with the Go binary
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
# COPY internal/db/migration ./migrations
EXPOSE 8080
CMD [ "/app/main" ]
