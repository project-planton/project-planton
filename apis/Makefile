.PHONY: build
build: lint fmt
	rm -rf generated/stubs generated/docs
	mkdir -p generated/stubs generated/docs
	buf generate
	cp -R generated/stubs/go/github.com/project-planton/project-planton/apis/. .
	cp -R generated/stubs/ts/org/project_planton/. ../app/frontend/src/gen/
	rm -rf generated/stubs

.PHONY: fmt
fmt:
	buf format -w

.PHONY: lint
lint: buf-lint

.PHONY: build-optional-linter-plugin
build-optional-linter-plugin:
	@echo "Building buf lint plugin..."
	@cd ../buf/lint/optional-linter && $(MAKE) build

.PHONY: buf-lint
buf-lint:
	buf lint
