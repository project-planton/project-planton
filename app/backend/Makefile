.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies downloaded"

.PHONY: generate
generate:
	@echo "Generating proto code from root apis/..."
	@cd ../../apis && buf generate 2>/dev/null || echo "⚠️  Warning: Proto generation had issues, but continuing..."
	@echo "Copying generated backend proto files..."
	@mkdir -p apis/gen/go/proto/backendv1connect
	@rm -f apis/gen/go/proto/*.pb.go apis/gen/go/proto/backendv1connect/*.connect.go
	@if [ -d "../../apis/generated/stubs/go/github.com/project-planton/project-planton/apis/org/project_planton/app" ]; then \
		for service_dir in ../../apis/generated/stubs/go/github.com/project-planton/project-planton/apis/org/project_planton/app/*/v1; do \
			if [ -d "$$service_dir" ]; then \
				find "$$service_dir" -name "*connect.go" -exec cp {} apis/gen/go/proto/backendv1connect/ \; 2>/dev/null || true; \
				find "$$service_dir" -name "*.pb.go" -exec cp {} apis/gen/go/proto/ \; 2>/dev/null || true; \
			fi; \
		done; \
		echo "✅ Backend proto files copied"; \
	else \
		echo "⚠️  Warning: Generated stubs not found. Run 'make protos' from project root first."; \
	fi

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
	@echo "✅ Clean complete"

.PHONY: test
test:
	@echo "Running tests..."
	go test ./...
	@echo "✅ Tests complete"
