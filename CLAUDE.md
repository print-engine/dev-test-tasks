# CLAUDE.md - AI-Assisted Development Guide

This document provides comprehensive guidance on using Claude Code and other AI tools to work with this dev-test-tasks repository effectively.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Repository Structure](#repository-structure)
3. [Working with Claude Code](#working-with-claude-code)
4. [Creating Test Tasks](#creating-test-tasks)
5. [Branch Naming Strategy](#branch-naming-strategy)
6. [AI-Assisted Development Best Practices](#ai-assisted-development-best-practices)
7. [Task Design Guidelines](#task-design-guidelines)
8. [Evaluation Criteria](#evaluation-criteria)

## Project Overview

This repository serves as a template and task collection for evaluating software engineering candidates across different experience levels. The main branch contains production-ready boilerplate code for both backend (Go) and frontend (React) projects.

### Purpose

The repository aims to assess three key areas:

1. **Technical Proficiency**: Candidate's experience with the tech stack and ability to write clean, maintainable code
2. **AI Fluency**: How effectively candidates leverage AI tools (like Claude Code, GitHub Copilot, etc.) to enhance productivity
3. **Problem-Solving Skills**: Ability to break down problems, make architectural decisions, and explain their thought process

## Repository Structure

```
dev-test-tasks/
├── backend/              # Go backend template
│   ├── main.go          # HTTP server with logging middleware
│   ├── main_test.go     # Comprehensive test suite
│   ├── go.mod           # Go module definition
│   └── README.md        # Backend documentation
├── frontend/            # React frontend template
│   ├── src/
│   │   ├── features/    # Redux features (example: counter)
│   │   ├── routes/      # React Router pages
│   │   ├── store/       # Redux store + sagas
│   │   └── test/        # Test configuration
│   ├── package.json     # Dependencies and scripts
│   ├── vitest.config.ts # Test configuration
│   └── README.md        # Frontend documentation
├── CLAUDE.md            # This file
├── README.md            # Project overview
└── .gitignore           # Git ignore rules
```

### Tech Stack

**Backend:**
- Go 1.25+
- Standard library HTTP server
- Production-ready patterns (middleware, logging, timeouts)
- Go native testing

**Frontend:**
- React 19 + TypeScript
- Vite for build tooling
- Redux Toolkit + Redux Saga
- React Router v7
- Vitest + React Testing Library
- ESLint + Prettier

## Working with Claude Code

### Initial Setup

When starting work on a task, Claude Code can help you:

1. **Understand the codebase**
   ```
   Can you explain the structure of the backend/frontend project?
   ```

2. **Set up development environment**
   ```
   Help me set up the development environment for the frontend task
   ```

3. **Run tests and verify setup**
   ```
   Run the tests and make sure everything is working
   ```

### Effective Prompting Strategies

#### ✅ Good Prompts

- **Specific and contextual**: "Add a new Redux slice for user authentication with login, logout, and token refresh actions"
- **Include requirements**: "Create an API endpoint for user registration that validates email format and password strength, returns appropriate HTTP status codes"
- **Request explanations**: "Explain the trade-offs between using Redux Saga vs Redux Thunk for this use case"

#### ❌ Less Effective Prompts

- **Too vague**: "Make it better"
- **No context**: "Add authentication" (without specifying requirements)
- **Overly broad**: "Build the entire feature"

### Iterative Development with Claude

Use an iterative approach:

1. **Plan first**: "Let's plan how to implement feature X. What files need to change?"
2. **Implement incrementally**: "Start with the data model, then add the API endpoint"
3. **Test along the way**: "Write tests for the new functionality"
4. **Refactor**: "Can we refactor this to be more maintainable?"

### Code Review with Claude

```
Review this implementation for:
- Security vulnerabilities
- Performance issues
- Best practices violations
- Test coverage gaps
```

## Creating Test Tasks

### Step 1: Create a New Branch

Follow the naming convention:

```bash
git checkout -b FE/middle/user-auth-flow
# or
git checkout -b BE/senior/api-rate-limiting
```

### Step 2: Define the Task

Create a `TASK.md` file in the branch with:

```markdown
# Task: [Task Name]

## Level: [Junior|Middle|Senior]
## Domain: [FE|BE]
## Estimated Time: [X hours]

## Objective
[Clear description of what needs to be built]

## Requirements
- [ ] Requirement 1
- [ ] Requirement 2
- [ ] Requirement 3

## Acceptance Criteria
- [ ] Criterion 1
- [ ] Criterion 2

## Evaluation Focus
- Technical implementation quality
- Code organization and patterns
- Test coverage
- AI tool usage efficiency

## Resources
- [Link to relevant documentation]
- [API specifications if needed]

## Deliverables
- Working implementation
- Tests
- Brief documentation of approach
```

### Step 3: Implement the Task

Starting from the main branch template, implement the task requirements. Ensure:

- Code builds successfully
- All tests pass
- Documentation is updated
- Task complexity matches the target level

### Step 4: Document AI Usage Patterns

In the task description, note:
- Which parts could benefit from AI assistance
- Where AI might struggle
- Expected workflow patterns

## Branch Naming Strategy

Format: `<Domain>/<Level>/<task-short-name>`

**Domain:**
- `FE` - Frontend (React)
- `BE` - Backend (Go)
- `FS` - Full-stack (both)

**Level:**
- `junior` - 0-2 years experience
- `middle` - 2-5 years experience
- `senior` - 5+ years experience

**Examples:**
- `FE/junior/todo-list-app`
- `BE/middle/rest-api-crud`
- `FE/senior/performance-optimization`
- `BE/senior/distributed-caching`
- `FS/middle/auth-integration`

## AI-Assisted Development Best Practices

### 1. Use AI for Boilerplate

AI excels at generating repetitive code:
- Test files
- Type definitions
- Configuration files
- Standard CRUD operations

### 2. Leverage AI for Documentation

- Generate comprehensive comments
- Create README files
- Write API documentation
- Explain complex algorithms

### 3. AI-Assisted Testing

- Generate test cases
- Create test data
- Write edge case tests
- Review test coverage

### 4. Code Review and Refactoring

- Identify code smells
- Suggest improvements
- Check for security issues
- Optimize performance

### 5. Learning and Research

- Explain unfamiliar concepts
- Compare different approaches
- Understand error messages
- Learn best practices

### 6. When NOT to Rely Solely on AI

- Critical security decisions
- Complex architectural choices
- Performance-critical code (verify AI suggestions)
- Business logic validation

## Task Design Guidelines

### Junior Level Tasks

**Focus:**
- Basic CRUD operations
- Simple state management
- Using existing patterns
- Following documentation

**Example Tasks:**
- Create a todo list component with add/remove functionality
- Implement a simple REST endpoint with validation
- Add a new page to the existing React app
- Write unit tests for a given function

**AI Usage Expected:**
- Heavy assistance with syntax
- Boilerplate generation
- Documentation lookup
- Debugging help

### Middle Level Tasks

**Focus:**
- Feature implementation with some architectural decisions
- Integrating multiple systems
- Error handling and edge cases
- Performance considerations

**Example Tasks:**
- Implement user authentication flow
- Create a reusable form system with validation
- Build a caching layer for API responses
- Implement pagination and filtering

**AI Usage Expected:**
- Pattern suggestions
- Implementation guidance
- Code review
- Testing strategies

### Senior Level Tasks

**Focus:**
- System design decisions
- Performance optimization
- Scalability considerations
- Complex business logic

**Example Tasks:**
- Design and implement a real-time notification system
- Optimize rendering performance for large datasets
- Implement a sophisticated caching strategy
- Design an extensible plugin architecture

**AI Usage Expected:**
- Architecture discussion
- Trade-off analysis
- Advanced pattern suggestions
- Security review

## Evaluation Criteria

### Technical Proficiency (40%)

- Code quality and organization
- Understanding of language/framework features
- Testing approach
- Error handling

### AI Fluency (30%)

- Effectiveness of AI tool usage
- Prompt quality and iteration
- Verification of AI-generated code
- Balance between AI assistance and own understanding

### Problem-Solving (30%)

- Approach to breaking down the problem
- Decision-making process
- Trade-off considerations
- Documentation of thought process

## Tips for Candidates Using Claude Code

1. **Start with exploration**: Ask Claude to explain the existing codebase
2. **Plan before coding**: Discuss the approach before implementing
3. **Test incrementally**: Run tests frequently with Claude's help
4. **Ask "why" questions**: Understand the reasoning behind suggestions
5. **Request alternatives**: "What are other ways to implement this?"
6. **Verify and validate**: Always review AI-generated code
7. **Document your process**: Keep track of your problem-solving approach

## Common Workflows

### Adding a New Feature

```bash
# 1. Understand the task
"Explain what I need to implement based on TASK.md"

# 2. Plan the implementation
"Let's plan the structure. What files do we need?"

# 3. Implement incrementally
"Start with the data model"
"Now add the API endpoint"
"Add the frontend component"

# 4. Test
"Write tests for this functionality"
"Run all tests and fix any failures"

# 5. Review
"Review this code for issues"
```

### Debugging

```bash
# 1. Identify the issue
"I'm getting this error: [error message]. What's wrong?"

# 2. Understand the cause
"Why would this cause the error?"

# 3. Fix
"How can we fix this?"

# 4. Prevent
"How can we prevent this type of error in the future?"
```

### Refactoring

```bash
# 1. Identify opportunities
"Review this code. What can be improved?"

# 2. Plan changes
"Let's refactor the authentication logic. What's the approach?"

# 3. Implement safely
"Refactor this while maintaining all existing tests"

# 4. Verify
"Run tests to ensure nothing broke"
```

## Conclusion

This repository is designed to evaluate both technical skills and the ability to effectively leverage AI tools in modern software development. The key is finding the right balance: using AI to enhance productivity while maintaining deep understanding of the code being written.

Remember: AI is a powerful tool, but the developer's judgment, understanding, and verification are irreplaceable.

---

**Questions or Issues?**

If you encounter any issues with this repository or have suggestions for improvements, please document them and discuss with your interviewer or repository maintainer.
