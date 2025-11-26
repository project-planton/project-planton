# Find buf command - check PATH first, then fallback to /tmp/buf
BUF := $(shell which buf 2>/dev/null || echo /tmp/buf)


# --------------------------------------------------------------------------- #
#  Go protobuf/Connect RPC stub generation
#  Following the same pattern as planton-cloud's Java generation
# --------------------------------------------------------------------------- #
.PHONY: go-stubs-clean
go-stubs-clean:
	@rm -rf gen/go

.PHONY: go-stubs-init
go-stubs-init:
	@mkdir -p gen/go

.PHONY: go-stubs
go-stubs: go-stubs-init
	@echo "Generating Go stubs from buf APIs..."
	@max_attempts=5; \
	attempt=1; \
	delay=2; \
	while [ $$attempt -le $$max_attempts ]; do \
		echo "Attempt $$attempt/$$max_attempts..."; \
		output=$$($(BUF) generate 2>&1); \
		exit_code=$$?; \
		if [ $$exit_code -eq 0 ]; then \
			echo "✅ Go stubs generated successfully"; \
			break; \
		else \
			if echo "$$output" | grep -q "resource_exhausted\|too many requests\|rate limit"; then \
				if [ $$attempt -lt $$max_attempts ]; then \
					echo "⚠️  Rate limit hit. Waiting $$delay seconds before retry..."; \
					sleep $$delay; \
					delay=$$((delay * 2)); \
					attempt=$$((attempt + 1)); \
				else \
					echo "❌ Rate limit error persists after $$max_attempts attempts."; \
					exit 1; \
				fi; \
			else \
				echo "$$output"; \
				echo "❌ Generation failed with error code $$exit_code"; \
				exit $$exit_code; \
			fi; \
		fi; \
	done

# --------------------------------------------------------------------------- #
#  TypeScript protobuf/Connect RPC stub generation
# --------------------------------------------------------------------------- #
.PHONY: ts-stubs
ts-stubs:
	@echo "Generating TypeScript stubs..."
	@./scripts/generate-ts-stubs.sh || true
	@echo "✅ TypeScript stubs generated"

# --------------------------------------------------------------------------- #
#  Main generation target
# --------------------------------------------------------------------------- #
.PHONY: generate
generate: go-stubs ts-stubs
	@echo "✅ All proto code generated"

.PHONY: lint
lint:
	@echo "Linting proto files..."
	@$(BUF) lint
	@echo "✅ Linting complete"

.PHONY: format
format:
	@echo "Formatting proto files..."
	@$(BUF) format -w
	@echo "✅ Formatting complete"

.PHONY: clean
clean: go-stubs-clean
	@echo "✅ Cleaned generated stubs"

.PHONY: build
build: lint format generate

