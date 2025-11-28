.PHONY: help build test run dev clean install lint

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the server binary
	@chmod +x scripts/build.sh
	@./scripts/build.sh

test: ## Run tests
	@chmod +x scripts/test.sh
	@./scripts/test.sh

test-coverage: ## Run tests with HTML coverage report
	@chmod +x scripts/test.sh
	@./scripts/test.sh --html

run: ## Run the server
	@go run cmd/server/main.go

dev: ## Run the server with hot reload
	@chmod +x scripts/dev.sh
	@./scripts/dev.sh

clean: ## Clean build artifacts
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

install: ## Install dependencies
	@go mod download
	@go mod tidy

lint: ## Run linters
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install: https://golangci-lint.run/usage/install/"; \
	fi

fmt: ## Format code
	@go fmt ./...
	@echo "Code formatted"

vet: ## Run go vet
	@go vet ./...
	@echo "Vet complete"

mod-verify: ## Verify dependencies
	@go mod verify
	@echo "Dependencies verified"

all: clean install fmt vet lint test build ## Run all checks and build
