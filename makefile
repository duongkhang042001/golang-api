# Makefile for Go project

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GORUN = $(GOCMD) run
GOTEST = $(GOCMD) test

# Binary output name
BINARY_NAME = core-api

# Main entry file
MAIN_FILE = cmd/api/main.go

# Build flags
BUILD_FLAGS =

# Build directory
BUILD_DIR = bin

# Default target
all: build run

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(BUILD_FLAGS) $(MAIN_FILE)
	@echo "Build complete."

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	@ ./$(BUILD_DIR)/$(BINARY_NAME)

# Clean the build
clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -f $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Cleaned."

# Run on mode develop the application
dev: 
	@echo "Starting development server..."
	@$(GORUN) $(MAIN_FILE)

.PHONY: all build run cleanS dev
