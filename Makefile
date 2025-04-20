.PHONY: build run test clean lint

# Application name
APP_NAME=supabase-e
BUILD_DIR=./build

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOLINT=golangci-lint

# Main application entry point
MAIN_PATH=./cmd/api

# Build the application
build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

# Run the application
run:
	$(GORUN) $(MAIN_PATH)/main.go

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Run linter
lint:
	$(GOLINT) run

# Install dependencies
deps:
	$(GOCMD) mod download

# Update dependencies
deps-update:
	$(GOCMD) get -u ./...
	$(GOCMD) mod tidy

# Help command
help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make lint          - Run linter"
	@echo "  make deps          - Install dependencies"
	@echo "  make deps-update   - Update dependencies"
