.PHONY: help dev build test clean install-deps docker-up docker-down migrate

# Default target
help:
	@echo "NautScanner - Available commands:"
	@echo "  make dev           - Start development environment"
	@echo "  make build         - Build backend and frontend"
	@echo "  make test          - Run tests"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make install-deps  - Install all dependencies"
	@echo "  make docker-up     - Start Docker containers"
	@echo "  make docker-down   - Stop Docker containers"
	@echo "  make migrate       - Run database migrations"

# Development
dev:
	@echo "Starting development environment..."
	docker-compose up -d postgres
	@echo "Waiting for database..."
	@sleep 3
	@echo "Starting backend..."
	cd backend && go run cmd/server/main.go &
	@echo "Starting frontend..."
	cd frontend && npm run dev

# Build
build:
	@echo "Building backend..."
	cd backend && go build -o ../bin/nautscanner cmd/server/main.go
	@echo "Building frontend..."
	cd frontend && npm run build

# Test
test:
	@echo "Running backend tests..."
	cd backend && go test ./...
	@echo "Running frontend tests..."
	cd frontend && npm test

# Clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -rf frontend/build/
	rm -rf frontend/.svelte-kit/

# Install dependencies
install-deps:
	@echo "Installing backend dependencies..."
	cd backend && go mod download
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

# Docker
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Database migrations
migrate:
	@echo "Running migrations..."
	cd backend && go run cmd/server/main.go migrate
