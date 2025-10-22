# Parallelism Project Makefile
.PHONY: all build run clean test bench fmt vet help init

# Project info
BINARY_NAME := parallelism
MODULE_NAME := parallelism

# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GORUN := $(GOCMD) run
GOFMT := $(GOCMD) fmt
GOVET := $(GOCMD) vet
GOMOD := $(GOCMD) mod

# Source files
GO_FILES := $(shell find . -name "*.go" -type f)
MAIN_FILE := main.go

# Colors for output
GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
NC := \033[0m # No Color

## Initialization
init:
	@echo "$(GREEN)Initializing project $(MODULE_NAME)...$(NC)"
	$(GOMOD) init $(MODULE_NAME)
	@echo "$(GREEN)Project initialized!$(NC)"

## Build targets
build:
	@echo "$(GREEN)Building $(BINARY_NAME)...$(NC)"
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

build-race:
	@echo "$(GREEN)Building with race detector...$(NC)"
	$(GOBUILD) -race -o $(BINARY_NAME)-race $(MAIN_FILE)

build-linux:
	@echo "$(GREEN)Building for Linux...$(NC)"
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-linux $(MAIN_FILE)

## Run targets
run: build
	@echo "$(GREEN)Running $(BINARY_NAME)...$(NC)"
	./$(BINARY_NAME)

run-dev:
	@echo "$(GREEN)Running in development mode...$(NC)"
	$(GORUN) $(MAIN_FILE)

run-race: build-race
	@echo "$(GREEN)Running with race detector...$(NC)"
	./$(BINARY_NAME)-race

## Test targets
test:
	@echo "$(GREEN)Running tests...$(NC)"
	$(GOTEST) -v ./...

test-race:
	@echo "$(GREEN)Running tests with race detector...$(NC)"
	$(GOTEST) -race -v ./...

test-coverage:
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	$(GOTEST) -cover -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "$(YELLOW)Coverage report: coverage.html$(NC)"

bench:
	@echo "$(GREEN)Running benchmarks...$(NC)"
	$(GOTEST) -bench=. -benchmem ./...

## Code quality
fmt:
	@echo "$(GREEN)Formatting code...$(NC)"
	$(GOFMT) ./...

fmt-check:
	@echo "$(GREEN)Checking code formatting...$(NC)"
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "$(RED)Code is not formatted:$(NC)"; \
		gofmt -l .; \
		exit 1; \
	else \
		echo "$(GREEN)Code is properly formatted$(NC)"; \
	fi

vet:
	@echo "$(GREEN)Vetting code...$(NC)"
	$(GOVET) ./...

lint: fmt-check vet
	@echo "$(GREEN)All linting checks passed!$(NC)"

## Dependencies
deps:
	@echo "$(GREEN)Downloading dependencies...$(NC)"
	$(GOMOD) download
	$(GOMOD) tidy

## Cleanup
clean:
	@echo "$(GREEN)Cleaning build artifacts...$(NC)"
	$(GOCLEAN)
	rm -f $(BINARY_NAME) $(BINARY_NAME)-race $(BINARY_NAME)-linux
	rm -f coverage.out coverage.html
	rm -f *.log

clean-all: clean
	@echo "$(GREEN)Cleaning all generated files...$(NC)"
	rm -rf dist/ bin/

## Development utilities
tree:
	@echo "$(GREEN)Project structure:$(NC)"
	@tree -I 'vendor|*.pb.go|coverage.*' .

check:
	@echo "$(GREEN)Checking project setup...$(NC)"
	@echo "Go version: $$(go version)"
	@echo "Module: $$(go list -m)"
	@echo "Files: $$(find . -name "*.go" | wc -l) Go files"
	@echo "Directories: $$(find . -type d | wc -l)"
	@echo "Import paths:"
	@echo "  - parallelism/libs"
	@echo "  - parallelism/crypto" 
	@echo "  - parallelism/types"

## Help
help:
	@echo "$(GREEN)Parallelism Project Makefile$(NC)"
	@echo ""
	@echo "$(YELLOW)Usage:$(NC)"
	@echo "  make [target]"
	@echo ""
	@echo "$(YELLOW)Targets:$(NC)"
	@echo "  $(GREEN)init$(NC)         - Initialize Go module"
	@echo "  $(GREEN)build$(NC)        - Build the binary (default)"
	@echo "  $(GREEN)build-race$(NC)   - Build with race detector"
	@echo "  $(GREEN)build-linux$(NC)  - Build for Linux"
	@echo ""
	@echo "  $(GREEN)run$(NC)          - Build and run"
	@echo "  $(GREEN)run-dev$(NC)      - Run without building"
	@echo "  $(GREEN)run-race$(NC)     - Run with race detector"
	@echo ""
	@echo "  $(GREEN)test$(NC)         - Run tests"
	@echo "  $(GREEN)test-race$(NC)    - Run tests with race detector"
	@echo "  $(GREEN)test-coverage$(NC)- Run tests with coverage report"
	@echo "  $(GREEN)bench$(NC)        - Run benchmarks"
	@echo ""
	@echo "  $(GREEN)fmt$(NC)          - Format code"
	@echo "  $(GREEN)fmt-check$(NC)    - Check code formatting"
	@echo "  $(GREEN)vet$(NC)          - Vet code"
	@echo "  $(GREEN)lint$(NC)         - Run all linting checks"
	@echo "  $(GREEN)deps$(NC)         - Download dependencies"
	@echo ""
	@echo "  $(GREEN)clean$(NC)        - Clean build artifacts"
	@echo "  $(GREEN)clean-all$(NC)    - Clean all generated files"
	@echo ""
	@echo "  $(GREEN)tree$(NC)         - Show project structure"
	@echo "  $(GREEN)check$(NC)        - Check project setup"
	@echo "  $(GREEN)help$(NC)         - Show this help message"

# Default target
.DEFAULT_GOAL := help