.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies downloaded"

.PHONY: generate
generate:
	@echo "Generating proto code..."
	cd apis && $(MAKE) generate
	@echo "✅ Code generation complete"

.PHONY: build
build: generate
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
	rm -rf apis/gen/
	@echo "✅ Clean complete"

.PHONY: test
test:
	@echo "Running tests..."
	go test ./...
	@echo "✅ Tests complete"

