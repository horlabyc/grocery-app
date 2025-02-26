FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o grocery-app ./cmd/api

# Create a lightweight production image
FROM alpine:3.18

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/grocery-app .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/.env.example ./.env

# Install migrate for database migrations
RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate && \
    apk del curl

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./grocery-app"]