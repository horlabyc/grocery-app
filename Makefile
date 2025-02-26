# Makefile

.PHONY: build run clean dev docker-dev docker-prod migrate-up migrate-down lint test

# Development
dev:
	air

# Docker development environment
docker-dev:
	docker-compose -f docker-compose.dev.yml up --build

# Docker production environment
docker-prod:
	docker-compose up --build -d

# Build the application
build:
	go build -o ./bin/grocery-app ./cmd/api

# Run the application
run: build
	./bin/grocery-app

# Clean build artifacts
clean:
	rm -rf ./bin ./tmp

# Database migrations
migrate-up:
	migrate -path=./migrations -database "postgres://admin:password@localhost:5432/os?sslmode=disable" up

migrate-down:
	migrate -path=./migrations -database "postgres://admin:password@localhost:5432/os?sslmode=disable" down

# Code quality
lint:
	golangci-lint run

# Testing
test:
	go test -v ./...