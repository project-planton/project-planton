.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies downloaded"

.PHONY: build
build:
	@echo "Building server..."
	go build -o bin/server ./cmd/server
	@echo "✅ Server built"

.PHONY: run
run: build
	@echo "Starting server..."
	./bin/server

.PHONY: dev
dev:
	@echo "Starting server in development mode..."
	go run ./cmd/server

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	@echo "✅ Clean complete"

.PHONY: test
test:
	@echo "Running tests..."
	go test ./...
	@echo "✅ Tests complete"
