# Backend - Go HTTP Server Template

A simple Go HTTP server template with logging middleware and health check endpoint.

## Features

- HTTP server with configurable port
- `/healthcheck` endpoint returning JSON status
- Logging middleware that tracks request method, path, duration, and status code
- Structured logging
- Comprehensive test coverage
- Production-ready server configuration (timeouts, graceful shutdown support)

## Prerequisites

- Go 1.25 or higher

## Getting Started

### Installation

No additional dependencies required - uses only Go standard library.

### Running the Server

```bash
# Run with default port (8080)
go run main.go

# Run with custom port
PORT=3000 go run main.go
```

### Building

```bash
# Build the binary
go build -o server

# Run the binary
./server
```

## API Endpoints

### Health Check

```
GET /healthcheck
```

Returns the server health status.

**Response:**
```json
{
  "status": "ok",
  "timestamp": "2025-11-14T15:55:00.123456Z"
}
```

**Status Code:** `200 OK`

## Testing

Run all tests:

```bash
go test -v
```

Run tests with coverage:

```bash
go test -v -cover
```

Generate coverage report:

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Project Structure

```
backend/
├── main.go           # Main server implementation
├── main_test.go      # Tests for server functionality
├── go.mod            # Go module definition
└── README.md         # This file
```

## Development

### Adding New Endpoints

1. Create a new handler function:
```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    // Your implementation
}
```

2. Register the handler in `main()`:
```go
mux.HandleFunc("/mypath", myHandler)
```

### Middleware

The logging middleware automatically wraps all registered handlers. To add additional middleware:

```go
handler := loggingMiddleware(mux, logger)
handler = myCustomMiddleware(handler)
```

## Configuration

The server can be configured using environment variables:

- `PORT`: Server port (default: 8080)

## Best Practices Demonstrated

- Middleware pattern for cross-cutting concerns
- Structured logging
- HTTP server timeouts for security
- JSON response formatting
- Comprehensive testing with table-driven tests
- Clean separation of concerns
