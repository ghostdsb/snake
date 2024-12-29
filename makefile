# Define variables
BINARY_NAME = SNAKE
SOURCE_FILES = $(wildcard *.go)

# Default target: Build the project
build: 
	@echo "Building the project..."
	go build -o $(BINARY_NAME) $(SOURCE_FILES)
	@echo "Build completed. Binary: $(BINARY_NAME)"

# Run the project
run: build
	@echo "Running the project..."
	./$(BINARY_NAME)

# Clean up build artifacts
clean: 
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	@echo "Cleanup completed."

# Run tests (if applicable)
test: 
	@echo "Running tests..."
	go test ./...
	@echo "Tests completed."

# Format the code
fmt:
	@echo "Formatting the code..."
	go fmt ./...
	@echo "Code formatted."

# Tidy up dependencies
tidy:
	@echo "Tidying up dependencies..."
	go mod tidy
	@echo "Dependencies tidied."

# Phony targets to avoid conflicts with actual files
.PHONY: build run clean test fmt tidy
