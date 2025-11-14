# Dev Test Tasks - Technical Assessment Repository

A comprehensive repository for evaluating software engineering candidates across different skill levels and domains. This repository serves as both a template and a collection of real-world coding tasks designed to assess technical proficiency, AI fluency, and problem-solving skills.

## Overview

This repository contains production-ready boilerplate code for both backend (Go) and frontend (React) projects. The main branch serves as a template, while feature branches contain specific test tasks targeting different experience levels.

## Strategic Goals

1. **Check Technical Proficiency**: Evaluate candidate's experience level and coding skills
2. **Assess AI Fluency**: Measure how effectively candidates leverage AI tools (Claude Code, GitHub Copilot, etc.)
3. **Reveal Problem-Solving**: Understand the candidate's thought process and decision-making approach

## Repository Structure

```
dev-test-tasks/
├── backend/              # Go backend template
│   ├── main.go          # HTTP server with /healthcheck endpoint
│   ├── main_test.go     # Comprehensive test suite
│   ├── go.mod           # Go module configuration
│   └── README.md        # Backend setup and usage guide
│
├── frontend/            # React frontend template
│   ├── src/
│   │   ├── features/    # Redux features (counter example)
│   │   ├── routes/      # React Router pages
│   │   ├── store/       # Redux store with saga middleware
│   │   └── test/        # Testing setup
│   ├── package.json     # Dependencies and scripts
│   └── README.md        # Frontend setup and usage guide
│
├── CLAUDE.md           # Comprehensive AI-assisted development guide
├── README.md           # This file
└── .gitignore          # Git ignore patterns
```

## Tech Stack

### Backend (Go + GCP)
- **Go 1.25+**: Modern Go with latest features
- **Standard Library**: HTTP server, logging, testing
- **Architecture**: RESTful API with middleware pattern
- **Testing**: Native Go testing with comprehensive coverage

**Features:**
- Production-ready HTTP server
- Health check endpoint (`/healthcheck`)
- Logging middleware (request tracking)
- Structured logging
- Comprehensive test coverage
- Security best practices (timeouts, proper error handling)

### Frontend (React + Firebase)
- **React 19**: Latest React with TypeScript
- **Vite**: Lightning-fast build tool
- **State Management**: Redux Toolkit + Redux Saga
- **Routing**: React Router v7
- **Testing**: Vitest + React Testing Library
- **Code Quality**: ESLint + Prettier
- **Type Safety**: TypeScript with strict mode

**Features:**
- Modern React architecture
- Type-safe Redux with hooks
- Declarative routing
- Comprehensive testing setup
- Code formatting and linting
- Example feature implementation (counter)

## Quick Start

### Backend

```bash
cd backend

# Run the server
go run main.go

# Run tests
go test -v

# Build
go build -o server
./server
```

**Server runs on**: `http://localhost:8080` (configurable via `PORT` env var)

**API Endpoints:**
- `GET /healthcheck` - Health status check

### Frontend

```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev

# Run tests
npm test

# Build for production
npm run build
```

**Development server**: `http://localhost:5173`

## Branch Naming Strategy

All test task branches follow this convention:

```
<Domain>/<Level>/<task-short-name>
```

**Domain:**
- `FE` - Frontend tasks (React)
- `BE` - Backend tasks (Go)
- `FS` - Full-stack tasks (both)

**Level:**
- `junior` - 0-2 years of experience
- `middle` - 2-5 years of experience
- `senior` - 5+ years of experience

**Examples:**
- `FE/junior/todo-list` - Simple frontend task for junior developers
- `BE/middle/api-caching` - Medium complexity backend task
- `FE/senior/performance-opt` - Advanced frontend optimization task
- `FS/middle/auth-flow` - Full-stack authentication implementation

## Creating New Test Tasks

### 1. Start from Main Branch

```bash
git checkout main
git pull origin main
```

### 2. Create a New Branch

```bash
git checkout -b FE/middle/user-dashboard
```

### 3. Implement the Task

- Build upon the existing template
- Add new functionality based on task requirements
- Ensure all tests pass
- Update documentation

### 4. Create Task Documentation

Add a `TASK.md` file describing:
- Objective and requirements
- Acceptance criteria
- Estimated time
- Evaluation focus areas
- Resources and hints

See `CLAUDE.md` for detailed guidelines on creating tasks.

## Evaluation Framework

### Technical Proficiency (40%)
- Code quality and organization
- Proper use of language/framework features
- Testing approach
- Error handling and edge cases
- Security considerations

### AI Fluency (30%)
- Effective use of AI coding assistants
- Quality of prompts and interactions
- Verification of AI-generated code
- Balance between AI assistance and understanding

### Problem-Solving Skills (30%)
- Approach to breaking down problems
- Decision-making and trade-offs
- Code documentation and comments
- Thought process clarity

## For Candidates

### Using This Repository

1. **Read the task description** in the branch-specific `TASK.md`
2. **Set up your environment** (see Quick Start above)
3. **Plan your approach** before coding
4. **Implement incrementally** with regular testing
5. **Document your decisions** and thought process
6. **Use AI tools effectively** (see CLAUDE.md for tips)

### AI Tool Usage

This repository is designed with AI-assisted development in mind. We encourage using tools like:
- Claude Code
- GitHub Copilot
- ChatGPT
- Other AI coding assistants

**Key Point**: We're evaluating how effectively you use these tools, not penalizing their use. The goal is to be productive while maintaining code understanding and quality.

See **CLAUDE.md** for comprehensive guidance on AI-assisted development patterns and best practices.

### What We're Looking For

- **Clean, maintainable code**: Easy to read and understand
- **Proper testing**: Unit and integration tests
- **Good architectural decisions**: Appropriate patterns and structure
- **Effective AI usage**: Leveraging AI to enhance productivity
- **Clear communication**: Documentation and code comments
- **Problem-solving approach**: How you break down and tackle problems

## Documentation

- **[CLAUDE.md](./CLAUDE.md)** - Comprehensive guide for AI-assisted development
- **[backend/README.md](./backend/README.md)** - Backend setup and development guide
- **[frontend/README.md](./frontend/README.md)** - Frontend setup and development guide

## Task Examples by Level

### Junior Level
- **FE**: Todo list with add/remove functionality
- **BE**: Simple CRUD REST API
- **Focus**: Following patterns, basic features, using documentation

### Middle Level
- **FE**: User authentication flow with form validation
- **BE**: API with caching and pagination
- **Focus**: Integrating systems, error handling, performance

### Senior Level
- **FE**: Performance optimization for large datasets
- **BE**: Distributed caching with fallback strategies
- **Focus**: Architecture, scalability, advanced patterns

## Prerequisites

### Backend Development
- Go 1.25 or higher
- Basic understanding of HTTP and REST APIs
- Familiarity with testing in Go

### Frontend Development
- Node.js 18+ and npm
- Understanding of React and TypeScript
- Experience with state management (Redux)
- Knowledge of modern frontend tooling

## Repository Maintenance

### Adding New Tasks

1. Create a branch following the naming convention
2. Implement the task starting from main template
3. Add comprehensive `TASK.md` documentation
4. Ensure all tests pass
5. Push the branch to remote

### Updating Templates

Changes to the main branch template should be:
- Well-tested
- Documented
- Backward compatible where possible
- Communicated to task maintainers

## License

This repository is for educational and assessment purposes.

## Questions or Issues?

For questions about:
- **Task requirements**: Check the branch-specific `TASK.md`
- **Setup issues**: See respective README files in `backend/` or `frontend/`
- **AI tool usage**: Refer to `CLAUDE.md`
- **General questions**: Contact your interviewer or repository maintainer

---

**Good luck with your assessment!**

Remember: We're looking for a combination of technical skills, problem-solving ability, and effective use of modern development tools including AI assistants. Focus on writing clean, maintainable code and documenting your thought process.
