.PHONY: build
build: lint fmt
	rm -rf generated/stubs generated/docs
	mkdir -p generated/stubs generated/docs
	buf generate
	cp -R generated/stubs/go/github.com/project-planton/project-planton/apis/. .
	rm -rf generated/stubs

.PHONY: fmt
fmt:
	buf format -w

.PHONY: lint
lint: build-lint-plugin buf-lint

.PHONY: build-lint-plugin
build-lint-plugin:
	@echo "Building buf lint plugin..."
	@cd ../buf/lint/planton && go build -o $(shell go env GOPATH)/bin/buf-plugin-planton ./cmd/buf-plugin-planton

.PHONY: buf-lint
buf-lint:
	buf lint
