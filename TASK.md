# Task: Real-Time WebSocket Notifications Service

## Overview

**Level:** Senior Backend Engineer  
**Domain:** Backend (Go)  
**Estimated Time:** 3-4 hours  
**Focus:** Concurrency, System Architecture, Error Handling, Testing

## Objective

Build a production-quality WebSocket notifications service that supports multiple concurrent clients, various messaging patterns, and message persistence. This task evaluates your ability to design and implement concurrent systems in Go while effectively leveraging AI development tools.

## Scenario

You're building a real-time notifications service that will handle multiple types of messages:
- Broadcast notifications (sent to all connected clients)
- Direct messages (sent to specific clients)
- Room-based messages (sent to clients in a specific room/channel)
- System events (user logins, alerts, errors)

The service needs to handle multiple concurrent connections, persist important messages, and gracefully handle network failures and client disconnections.

## Functional Requirements

### Core Requirements (Must Have)

1. **WebSocket Connection Management**
   - Accept multiple concurrent WebSocket connections
   - Assign each client a unique identifier
   - Maintain a registry of active connections
   - Handle client disconnections gracefully
   - Clean up resources when clients disconnect

2. **Broadcast Messaging**
   - Accept messages from any connected client
   - Broadcast messages to all other connected clients
   - Message format: JSON with type and content
   - Example: `{"type": "broadcast", "message": "Hello everyone"}`

3. **Direct Messaging**
   - Send messages to specific clients by their ID
   - Return error if target client doesn't exist or is disconnected
   - Message format: `{"type": "direct", "targetId": "client-123", "message": "Hello"}`

4. **Room-Based Messaging**
   - Clients can join named rooms/channels
   - Clients can leave rooms
   - Messages sent to a room only reach clients in that room
   - Clients can be in multiple rooms simultaneously
   - Message format: `{"type": "join|leave|room_message", "room": "general", "message": "..."}`

5. **Event Handling**
   - Process different event types (user_login, system_alert, error, custom)
   - Broadcast events to appropriate recipients
   - Event format: `{"type": "event", "eventType": "user_login", "data": {...}}`

6. **Message Persistence**
   - Store all messages in the database
   - Load recent message history when clients connect
   - Include metadata: timestamp, sender, recipients, message type
   - Implement reasonable retention (e.g., last 100 messages or 24 hours)

## Non-Functional Requirements

### Concurrency & Performance
- Handle at least 100 concurrent connections without crashing
- All shared state must be thread-safe (no race conditions)
- Efficient message routing (avoid unnecessary iterations)
- No goroutine leaks
- Proper use of channels for inter-goroutine communication

### Error Handling & Resilience
- Graceful handling of:
  - Invalid message formats
  - Network errors
  - Client disconnections mid-message
  - Database write failures
- Comprehensive error logging
- Don't crash on malformed input
- Recover from panics in goroutines

### Code Quality
- Clean, readable, and maintainable code
- Proper separation of concerns
- Appropriate use of Go idioms and patterns
- Comprehensive inline documentation where complexity exists
- No hardcoded values (use configuration)

### Testing
- Unit tests for core logic
- Integration tests for WebSocket handlers
- Test concurrent scenarios (race detector compatible)
- Test error handling paths
- Minimum 70% code coverage for critical paths

## Acceptance Criteria

Use the provided frontend client (`docs/index.html`) to verify:

- [ ] **Multiple clients can connect simultaneously**
  - Open 3+ browser windows
  - All should connect successfully
  - Each receives a unique client ID

- [ ] **Broadcast messages work correctly**
  - Client A sends a broadcast message
  - All other clients (B, C, D...) receive it
  - Sender does not receive their own message

- [ ] **Direct messages reach only the target**
  - Client A sends direct message to Client B's ID
  - Only Client B receives the message
  - Other clients don't see it
  - Error handling when target doesn't exist

- [ ] **Room functionality works**
  - Clients can join a room (e.g., "general")
  - Messages to that room reach only room members
  - Clients not in the room don't receive those messages
  - Clients can leave rooms
  - Clients can be in multiple rooms

- [ ] **Message persistence works**
  - Send several messages
  - Disconnect and reconnect a client
  - Client receives recent message history on reconnection

- [ ] **Event handling works**
  - Click event simulator buttons in frontend
  - Events are properly broadcast or routed
  - Event structure is preserved

- [ ] **Stress test passes**
  - Use the "Connection Stress Test" feature
  - Create 100 rapid connections
  - Server remains stable
  - All connections establish successfully or fail gracefully

- [ ] **Graceful disconnect handling**
  - Disconnect a client (close browser tab)
  - Other clients continue functioning
  - No errors in server logs
  - Resources cleaned up

- [ ] **Tests pass**
  - Run `go test ./...` - all tests pass
  - Run `go test -race ./...` - no race conditions
  - Reasonable test coverage of critical paths

- [ ] **Code quality**
  - Clean, organized code structure
  - Appropriate use of Go concurrency primitives
  - Error handling throughout
  - No obvious security issues

## Message Format Specification

### Incoming Messages (Client → Server)

```json
// Broadcast
{
  "type": "broadcast",
  "message": "string"
}

// Direct message
{
  "type": "direct",
  "targetId": "client-uuid",
  "message": "string"
}

// Join room
{
  "type": "join",
  "room": "room-name"
}

// Leave room
{
  "type": "leave",
  "room": "room-name"
}

// Room message
{
  "type": "room_message",
  "room": "room-name",
  "message": "string"
}

// Event
{
  "type": "event",
  "eventType": "user_login|system_alert|error|custom",
  "data": { /* event-specific data */ }
}
```

### Outgoing Messages (Server → Client)

```json
// Regular message
{
  "from": "client-uuid",
  "type": "broadcast|direct|room_message|event",
  "message": "string",
  "timestamp": "2024-01-01T12:00:00Z"
}

// System message (connection, error, etc.)
{
  "type": "system",
  "message": "string",
  "data": { /* optional additional data */ }
}

// Initial connection
{
  "type": "connected",
  "clientId": "your-client-uuid",
  "messageHistory": [ /* recent messages */ ]
}

// Error
{
  "type": "error",
  "message": "error description"
}
```

## Evaluation Focus

This assessment evaluates:

### Technical Implementation (60%)
- **Concurrency patterns:** Proper use of goroutines, channels, mutexes
- **Architecture:** Clean separation of concerns, appropriate abstractions
- **Error handling:** Comprehensive error handling and recovery
- **Testing:** Quality and coverage of tests

### AI Collaboration (40%)
- **Prompt quality:** How effectively you communicate with the AI assistant
- **Problem decomposition:** Breaking down the task into manageable pieces
- **Verification:** Validating AI-generated code and suggestions
- **Iteration:** Refining and improving with AI assistance
- **Understanding:** Demonstrating comprehension of the implemented solution

## Constraints & Guidelines

- Use the existing project structure in `backend/`
- The `gorilla/websocket` library is already added to dependencies
- Use SQLite for persistence (already configured)
- Build upon provided infrastructure (logging, config, middleware)
- Don't modify the frontend client code
- Feel free to add new packages and files as needed
- Document any assumptions you make

## Stretch Goals (Optional)

If you complete the core requirements early, consider:

- **Ping/Pong mechanism:** Detect and clean up dead connections
- **Rate limiting:** Prevent clients from flooding the server
- **Authentication:** Basic token-based client authentication
- **Message queuing:** Handle slow consumers without blocking fast ones
- **Metrics:** Track connection count, message rate, errors
- **Admin API:** HTTP endpoints to view server state, connected clients
- **Graceful shutdown:** Properly close all connections on SIGINT/SIGTERM

## Tips

- Start simple: Get basic broadcast working first, then add complexity
- Test frequently using the provided frontend client
- Use the race detector (`go test -race`) early and often
- Log important events to help with debugging
- Think about failure modes: what happens when connections drop?
- Ask clarifying questions if requirements are ambiguous
- Document your architectural decisions
- Collaborate effectively with your AI assistant

## Resources

- `CLAUDE.md` - Domain context and patterns
- `README.md` - Project structure and how to run
- `docs/index.html` - Frontend testing client
- Gorilla WebSocket docs: https://pkg.go.dev/github.com/gorilla/websocket

---

**Remember:** This is a time-boxed assessment. A working implementation of core features with good architecture is better than incomplete advanced features. Focus on correctness, clarity, and demonstrating your thought process.

**Good luck!**
