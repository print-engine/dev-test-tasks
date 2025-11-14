# Backend - Go HTTP Server with Clean Architecture

A production-ready Go HTTP server template following Clean Architecture principles, featuring comprehensive project structure, configuration management, and development tooling.

## Features

- **Clean Architecture**: Organized with clear separation of concerns (cmd, internal, pkg, config)
- **HTTP Server**: Production-ready server with configurable timeouts
- **Health Check API**: `/healthcheck` endpoint with JSON response
- **Logging Middleware**: Request tracking with method, path, duration, and status code
- **Configuration Management**: Environment-based config with sensible defaults
- **Database Migrations**: Structure and examples for schema management
- **Build Scripts**: Automated build, test, and development workflows
- **API Documentation**: OpenAPI 3.0 specification
- **Comprehensive Tests**: Unit tests for all components

## Prerequisites

- Go 1.25 or higher
- Make (optional, for using Makefile commands)

## Project Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/                     # Private application code
│   ├── handlers/                 # HTTP handlers
│   │   ├── health.go
│   │   └── health_test.go
│   ├── middleware/               # HTTP middleware
│   │   ├── logging.go
│   │   └── logging_test.go
│   ├── models/                   # Domain models
│   │   └── health.go
│   ├── service/                  # Business logic (empty, for extension)
│   └── repository/               # Data access (empty, for extension)
├── pkg/                          # Public, reusable packages
│   └── logger/
│       ├── logger.go
│       └── logger_test.go
├── config/                       # Configuration management
│   ├── config.go
│   └── .env.example
├── api/                          # API specifications
│   └── openapi.yaml              # OpenAPI 3.0 spec
├── migrations/                   # Database migrations
│   ├── README.md
│   ├── 000001_create_users_table.up.sql
│   └── 000001_create_users_table.down.sql
├── scripts/                      # Build and deployment scripts
│   ├── build.sh
│   ├── test.sh
│   └── dev.sh
├── .air.toml                     # Hot reload configuration
├── Makefile                      # Development commands
├── go.mod                        # Go module definition
└── README.md                     # This file
```

## Quick Start

### Using Makefile (Recommended)

```bash
# See all available commands
make help

# Run the server
make run

# Run with hot reload
make dev

# Run tests
make test

# Build binary
make build
```

### Manual Commands

```bash
# Install dependencies
go mod download

# Run the server
go run cmd/server/main.go

# Run with custom port
PORT=3000 go run cmd/server/main.go

# Run tests
go test ./...

# Build the binary
go build -o bin/server cmd/server/main.go
```

## Configuration

The server is configured via environment variables. See `config/.env.example` for all available options.

### Server Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `READ_TIMEOUT` | `10s` | HTTP read timeout |
| `WRITE_TIMEOUT` | `10s` | HTTP write timeout |
| `IDLE_TIMEOUT` | `60s` | HTTP idle timeout |

### Database Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | `` | Database password |
| `DB_NAME` | `testdb` | Database name |
| `DB_SSLMODE` | `disable` | SSL mode |

## API Documentation

The API is documented using OpenAPI 3.0. See `api/openapi.yaml` for the full specification.

### Endpoints

#### Health Check

```
GET /healthcheck
```

Returns the server health status.

**Response (200 OK):**
```json
{
  "status": "ok",
  "timestamp": "2025-11-14T15:55:00.123456Z"
}
```

## Development

### Project Organization

**cmd/server/**: Application entry point. Wires together all components.

**internal/**: Private application code that cannot be imported by external projects.
- `handlers/`: HTTP request handlers
- `middleware/`: HTTP middleware (logging, auth, etc.)
- `models/`: Domain models and DTOs
- `service/`: Business logic layer
- `repository/`: Data access layer

**pkg/**: Public packages that can be imported by external projects.

**config/**: Configuration management and environment variable handling.

**api/**: API specifications and documentation.

**migrations/**: Database migration files.

**scripts/**: Build, test, and deployment automation.

### Adding New Features

#### 1. Adding a New Endpoint

**Step 1**: Create a handler in `internal/handlers/`:

```go
// internal/handlers/users.go
package handlers

import (
    "net/http"
    "dev-test-tasks/internal/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    // Implementation
}
```

**Step 2**: Register the route in `cmd/server/main.go`:

```go
mux.HandleFunc("/users/{id}", handlers.GetUser)
```

**Step 3**: Add tests in `internal/handlers/users_test.go`

#### 2. Adding Middleware

Create middleware in `internal/middleware/`:

```go
// internal/middleware/auth.go
package middleware

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Auth logic
        next.ServeHTTP(w, r)
    })
}
```

Apply in `cmd/server/main.go`:

```go
handler := middleware.Logging(mux, log)
handler = middleware.Auth(handler)
```

#### 3. Adding Configuration

Update `config/config.go` to add new config fields:

```go
type Config struct {
    Server ServerConfig
    DB     DatabaseConfig
    NewFeature NewFeatureConfig  // Add new section
}
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run tests for specific package
go test ./internal/handlers/...

# Run with race detector
go test -race ./...
```

### Building

```bash
# Build with Makefile
make build

# Build manually
go build -o bin/server cmd/server/main.go

# Build with version info
VERSION=1.0.0 make build
```

### Hot Reload Development

```bash
# Install air if not present and run
make dev

# Or manually
air -c .air.toml
```

### Code Quality

```bash
# Format code
make fmt

# Run go vet
make vet

# Run linter (requires golangci-lint)
make lint

# Run all checks
make all
```

## Database Migrations

See `migrations/README.md` for detailed migration instructions.

Quick start with golang-migrate:

```bash
# Create new migration
migrate create -ext sql -dir migrations -seq add_posts_table

# Run migrations
migrate -path migrations -database "postgres://user:pass@localhost:5432/db?sslmode=disable" up

# Rollback
migrate -path migrations -database "postgres://user:pass@localhost:5432/db?sslmode=disable" down 1
```

## Architecture Guidelines

### Dependency Direction

Dependencies should flow inward:
```
cmd → internal → pkg
```

- `cmd`: Depends on internal and pkg
- `internal`: Depends on pkg, not on cmd
- `pkg`: Self-contained, no project dependencies

### Package Responsibilities

- **handlers**: HTTP request/response handling only
- **service**: Business logic, orchestration
- **repository**: Data access, database operations
- **models**: Data structures, no logic
- **middleware**: Cross-cutting concerns (logging, auth, etc.)
- **pkg**: Reusable utilities

### Best Practices

1. **Keep handlers thin**: Delegate to service layer
2. **Business logic in services**: Not in handlers or repositories
3. **Use interfaces**: For testability and flexibility
4. **Error handling**: Return errors, don't panic
5. **Context usage**: Pass context for cancellation and timeouts
6. **Logging**: Use structured logging
7. **Testing**: Write tests for all packages

## Production Deployment

### Building for Production

```bash
# Build optimized binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o bin/server \
    cmd/server/main.go
```

### Docker (Example)

```dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server cmd/server/main.go

FROM alpine:latest
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
```

### Environment Variables

Set production environment variables:

```bash
export PORT=8080
export DB_HOST=production-db.example.com
export DB_PASSWORD=secure_password
# ... other vars
```

## Troubleshooting

### Server won't start

- Check if port is already in use: `lsof -i :8080`
- Verify environment variables are set correctly
- Check logs for error messages

### Tests failing

- Run `go mod tidy` to ensure dependencies are correct
- Check if you're in the backend directory
- Verify Go version: `go version`

### Import errors

- Run `go mod download`
- Verify module path in `go.mod` matches your directory structure

## License

This is a template project for educational purposes.

## Contributing

This is a template repository for test tasks. Contributions should be made via creating new test task branches following the branch naming convention.
