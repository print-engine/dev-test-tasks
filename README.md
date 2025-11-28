# WebSocket Notifications Service - Senior Backend Assessment

This branch contains a technical assessment for Senior Backend Engineers focusing on building a real-time WebSocket notifications service in Go.

## Initial Setup (Important!)

Before you begin, you need to configure your environment with the Anthropic API key for Claude Code assistance.

### macOS / Linux

Run this command in your terminal:

```bash
source setup.sh
```

### Windows

**PowerShell (Recommended):**
```powershell
.\setup.ps1
```
_Works immediately in the current session!_

**Command Prompt:**
```cmd
setup.bat
```
_Note: You'll need to open a NEW Command Prompt window after running this._

The script will:
1. Ask you to paste the API key (provided by your interviewer)
2. Configure your environment
3. Give you next steps

**Then simply type `claude` to start!**

## What's Provided

### Backend Structure
The backend project is organized as follows:

```
backend/
├── cmd/server/          # Application entry point
│   └── main.go         # Server setup and routing
├── internal/           # Internal application packages
│   ├── handlers/       # HTTP and WebSocket handlers
│   ├── middleware/     # HTTP middleware (logging, CORS)
│   ├── models/         # Data models
│   ├── service/        # Business logic layer
│   └── repository/     # Data access layer
├── pkg/                # Public reusable packages
│   └── logger/         # Logging utilities
├── config/             # Configuration management
├── api/                # API specifications
├── migrations/         # Database migrations
├── data/               # SQLite database storage
└── scripts/            # Build and test automation scripts
```

### What Exists

**Server Infrastructure:**
- HTTP server with routing (`/healthcheck`, `/ws` endpoints)
- Logging middleware for request tracking
- CORS middleware configured for local development
- Basic configuration management via environment variables
- SQLite database support with migration system

**WebSocket Endpoint:**
- Basic `/ws` endpoint configured in routing
- Gorilla WebSocket library dependency added
- Minimal handler that performs WebSocket upgrade

**Testing Frontend:**
A fully functional WebSocket client is provided in `docs/index.html`:
- Connection management (connect/disconnect)
- Message sending (broadcast and direct)
- Room management (join/leave rooms)
- Event simulator (various event types)
- Connection stress testing
- Real-time message display

### Getting Started

#### Prerequisites
- Go 1.25 or higher
- Web browser (for testing with the frontend client)

#### Running the Backend Server

From the `backend` directory:

```bash
# Install dependencies
go mod download

# Run the server
go run cmd/server/main.go

# Or use the Makefile
make run

# Run in development mode with hot-reload
make dev
```

The server will start on `http://localhost:8080` (configurable via `PORT` environment variable).

#### Testing with the Frontend Client

1. Open `docs/index.html` in your web browser
2. The default WebSocket URL is `ws://localhost:8080/ws`
3. Click "Connect" to establish a WebSocket connection
4. Use the various controls to interact with your implementation

You can open multiple browser tabs/windows to simulate multiple clients.

#### Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with race detector
go test -race ./...

# Run tests with coverage
go test -cover ./...
```

### Available Make Commands

```bash
make build          # Build the application
make run            # Run the application
make dev            # Run with hot-reload (requires air)
make test           # Run tests
make test-coverage  # Run tests with coverage
make clean          # Clean build artifacts
make lint           # Run linter
make fmt            # Format code
```

## Task Requirements

See **[TASK.md](./TASK.md)** for:
- Detailed requirements and objectives
- Functional and non-functional requirements
- Acceptance criteria
- Time estimates
- Evaluation focus areas

## AI Assistant Guidance

See **[CLAUDE.md](./CLAUDE.md)** for:
- Problem domain context (WebSocket, Go concurrency)
- Testing strategies
- Common challenges and patterns
- Evaluation expectations
- Resources and documentation links

## Project Dependencies

### Go Modules
- `github.com/gorilla/websocket` - WebSocket protocol implementation
- `github.com/joho/godotenv` - Environment variable management
- Standard library packages for HTTP, testing, database

### Database
- SQLite3 for message persistence
- Migration support via `migrations/` directory

## Configuration

The server can be configured via environment variables or a `.env` file:

```bash
# Server settings
PORT=8080
READ_TIMEOUT=10s
WRITE_TIMEOUT=10s
IDLE_TIMEOUT=60s

# Database settings
DB_PATH=./data/database.db
```

See `config/config.go` and `.env.example` for all available configuration options.

## Development Workflow

1. **Understand the Requirements**: Read `TASK.md` carefully
2. **Plan Your Architecture**: Think about structure and patterns
3. **Implement Incrementally**: Build and test in small steps
4. **Test Continuously**: Use the frontend client and automated tests
5. **Iterate and Refine**: Improve based on testing and feedback

## Testing Your Implementation

### Manual Testing
1. Start the backend server
2. Open the frontend client in multiple browser windows
3. Test each feature:
   - Connect multiple clients
   - Send broadcast messages
   - Send direct messages to specific clients
   - Create and join rooms
   - Send room-scoped messages
   - Trigger event simulations
   - Run the stress test
   - Verify message persistence

### Automated Testing
Write tests covering:
- WebSocket connection handling
- Message routing logic
- Concurrent client management
- Room management
- Message persistence
- Error handling and edge cases

## Project Structure Patterns

This codebase follows Go best practices:

- **`cmd/`**: Application entry points
- **`internal/`**: Private application code
- **`pkg/`**: Public, reusable packages
- **`config/`**: Configuration management
- **`migrations/`**: Database schema changes
- **`data/`**: Runtime data storage

Handlers should focus on HTTP/WebSocket concerns, business logic belongs in the `service/` layer, and data access in the `repository/` layer.

## Debugging

### Server Logs
The server logs all requests via the logging middleware. Watch the console output for:
- Connection events
- Message routing
- Errors and warnings

### Frontend Client
The frontend displays all WebSocket events:
- Connection status changes
- Sent messages
- Received messages
- Errors

## Common Issues

**Port already in use:**
```bash
# Check what's using port 8080
lsof -i :8080

# Or change the port
PORT=3000 go run cmd/server/main.go
```

**WebSocket connection fails:**
- Ensure the server is running
- Check the WebSocket URL matches the server port
- Verify CORS middleware is configured correctly

**Tests failing:**
- Run with `-v` flag for detailed output
- Use race detector to find concurrency issues
- Check test isolation (each test should be independent)

## Resources

### Go Documentation
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)

### WebSocket
- [Gorilla WebSocket Documentation](https://pkg.go.dev/github.com/gorilla/websocket)
- [WebSocket Protocol RFC](https://tools.ietf.org/html/rfc6455)

### Testing
- [Go Testing Package](https://pkg.go.dev/testing)
- [Table-Driven Tests in Go](https://go.dev/wiki/TableDrivenTests)
- [Testing WebSockets](https://pkg.go.dev/github.com/gorilla/websocket#pkg-examples)

---

**Important**: This is a senior-level assessment. You are expected to architect your own solution based on the requirements. The provided structure is a foundation - how you build upon it is part of the evaluation.

Focus on:
- Correctness and reliability
- Clean, maintainable code
- Thoughtful architectural decisions
- Comprehensive testing
- Effective collaboration with AI tools

Good luck!
