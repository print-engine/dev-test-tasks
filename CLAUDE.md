# Senior Backend Developer Assessment - WebSocket Notifications Service

## Context for AI Assistants

This document provides domain context and evaluation criteria for this technical assessment. It is designed to help AI assistants understand the problem domain without providing specific implementation guidance.

## Assessment Overview

This is a **Senior-level Backend Engineering assessment** focusing on building a real-time WebSocket notifications service in Go. The assessment evaluates both **technical proficiency** and **AI collaboration effectiveness**.

### Dual Evaluation Criteria

**Technical Skills (60%)**
- Go concurrency patterns and best practices
- System architecture and design decisions
- Error handling and resilience
- Testing strategy and coverage

**AI Fluency (40%)**
- Quality and clarity of prompts to AI assistant
- Ability to break down complex problems
- Critical evaluation of AI-generated code
- Effective iteration and refinement with AI

## Problem Domain: Real-Time WebSocket Communication

### What is WebSocket?

WebSocket is a communication protocol providing full-duplex communication channels over a single TCP connection. Unlike HTTP's request-response model, WebSocket enables:
- Persistent, long-lived connections
- Bi-directional communication (client ↔ server)
- Low-latency message exchange
- Efficient real-time data transfer

### Common Use Cases
- Real-time notifications and alerts
- Live chat applications
- Collaborative editing tools
- Live dashboards and monitoring
- Multiplayer games
- Financial tickers and trading platforms

### WebSocket Lifecycle
1. **Handshake**: HTTP upgrade request
2. **Connection**: Persistent TCP connection established
3. **Messaging**: Bi-directional frame exchange
4. **Termination**: Graceful close or error disconnection

## Go Concurrency Fundamentals

### Core Concepts

**Goroutines**: Lightweight threads managed by Go runtime
- Cheap to create (thousands can run simultaneously)
- Scheduled cooperatively
- Share memory space (requires synchronization)

**Channels**: Type-safe communication between goroutines
- Unbuffered: Synchronous (blocking send/receive)
- Buffered: Asynchronous until buffer is full

**Select Statement**: Multiplexing channel operations
- Handle multiple channels
- Implement timeouts
- Non-blocking operations

### Synchronization Primitives

**Mutex (sync.Mutex)**: Mutual exclusion lock
- Protects shared data from concurrent access
- Lock before read/write, unlock after

**RWMutex (sync.RWMutex)**: Reader-writer lock
- Multiple readers OR single writer
- Better performance for read-heavy workloads

**WaitGroup (sync.WaitGroup)**: Wait for goroutine completion
- Counter-based synchronization
- Wait for multiple operations to complete

**Context (context.Context)**: Cancellation and timeout propagation
- Signal cancellation across goroutine boundaries
- Carry request-scoped values
- Implement timeouts and deadlines

## Architectural Considerations

### Scalability Patterns

**Vertical Scaling**
- More resources (CPU/memory) on single server
- Limited by hardware constraints
- Simpler to implement initially

**Horizontal Scaling**
- Multiple server instances
- Requires message routing/broadcasting across instances
- Common solutions: Redis Pub/Sub, Message Queues (NATS, RabbitMQ)

### Design Patterns for WebSocket Services

These are **general industry patterns** (not specific implementation instructions):

**Hub Pattern**: Central coordinator managing connections
**Client Registry**: Tracking active connections
**Message Router**: Directing messages to appropriate destinations
**Backpressure Handling**: Managing slow/unresponsive clients
**Graceful Shutdown**: Cleanup on server termination

### Common Challenges

**Connection Management**
- Tracking active connections
- Handling disconnections gracefully
- Detecting dead connections (ping/pong)

**Concurrency Safety**
- Protecting shared state
- Avoiding race conditions
- Preventing deadlocks

**Resource Management**
- Memory leaks from abandoned connections
- Goroutine leaks
- File descriptor limits

**Error Handling**
- Network failures
- Invalid message formats
- Client misbehavior

## Testing Concurrent Systems

### Testing Challenges
- Non-deterministic execution order
- Race conditions that appear intermittently
- Deadlocks under specific conditions

### Testing Strategies
- Unit tests for individual components
- Integration tests for end-to-end flows
- Race detector (`go test -race`)
- Load/stress testing for concurrency issues
- Mocking network connections

### Go Testing Tools
- `testing` package (standard library)
- `httptest` for HTTP/WebSocket testing
- `testify` for assertions (third-party)
- Race detector built into Go toolchain

## Performance Considerations

### Bottlenecks in WebSocket Services
- Lock contention (too coarse-grained locking)
- Channel blocking (slow consumers)
- Memory allocation (frequent allocations)
- System calls (context switching overhead)

### Optimization Techniques
- Fine-grained locking (reduce lock scope)
- Buffered channels (reduce blocking)
- Object pooling (reduce GC pressure)
- Goroutine pools (limit concurrency)

## Security Considerations

### Common Vulnerabilities
- Missing origin validation (CORS)
- Lack of authentication/authorization
- Message injection attacks
- Denial of Service (resource exhaustion)
- Sensitive data in logs

### Best Practices
- Validate all inputs
- Implement rate limiting
- Use TLS/WSS in production
- Sanitize log output
- Implement connection timeouts

## Database Integration

### Persistence Patterns
- Write-through: Write to DB immediately
- Write-behind: Queue writes, batch process
- Event sourcing: Store events, rebuild state

### Considerations
- Synchronous vs asynchronous writes
- Transaction boundaries
- Connection pooling
- Query performance

## Evaluation Expectations

### What We're Looking For

**Architecture**
- Clear separation of concerns
- Appropriate use of design patterns
- Consideration for scalability

**Concurrency**
- Correct use of goroutines and channels
- Proper synchronization
- No race conditions or deadlocks

**Error Handling**
- Graceful degradation
- Proper error propagation
- Resource cleanup

**Testing**
- Coverage of critical paths
- Testing concurrent behavior
- Clear test cases

**AI Collaboration**
- Thoughtful questions to AI
- Validation of AI suggestions
- Iterative refinement
- Understanding limitations

### What We're NOT Looking For
- Production-ready, enterprise-scale system (time-boxed assessment)
- Perfect optimization (premature optimization is harmful)
- Exhaustive test coverage (focus on critical paths)
- Complex patterns without justification (YAGNI principle)

## Resources

### Go Documentation
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Context Package](https://pkg.go.dev/context)

### WebSocket
- [RFC 6455](https://tools.ietf.org/html/rfc6455) - The WebSocket Protocol
- [Gorilla WebSocket](https://github.com/gorilla/websocket) - Go WebSocket library

### Testing
- [Go Testing](https://pkg.go.dev/testing)
- [Race Detector](https://go.dev/doc/articles/race_detector)

## Important Notes

- This is a **senior-level assessment** - you are expected to architect your own solution
- There are multiple valid approaches - choose what makes sense for the requirements
- Ask clarifying questions when requirements are ambiguous
- Focus on correctness first, optimization second
- Document your architectural decisions and trade-offs

## Time Management

- Estimated time: 3-4 hours
- Focus on core requirements first
- Implement stretch goals if time permits
- It's better to have a working MVP than incomplete advanced features

---

**Remember**: This assessment evaluates your ability to design, implement, and explain a concurrent system while effectively collaborating with AI tools. Show your thought process, ask good questions, and build something that works.
