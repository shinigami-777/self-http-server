# Variables
# TODO : Make the makefile.
APP_NAME := app
VERSION := 1.0.0
BUILD_DIR := build

# Default Target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	cd $(APP_NAME)
	go build

# Run the project
.PHONY: run
run:
	go run app/server.go

# Test the project
.PHONY: test
test:
	go test ./...

# Clean up build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Format the code
.PHONY: format
format:
	go fmt ./...

# Install dependencies
.PHONY: deps
deps:
	go mod tidy

# Lint the code
.PHONY: lint
lint:
	golangci-lint run


