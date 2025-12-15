.PHONY: build
build: lint fmt
	rm -rf generated/stubs
	mkdir -p generated/stubs
	buf generate --include-imports
	cp -R generated/stubs/go/github.com/project-planton/project-planton/apis/. .
	cp -R generated/stubs/ts/org/project_planton/. ../app/frontend/src/gen/
	# Copy buf/validate files if they exist (needed for validation imports)
	@if [ -d "generated/stubs/ts/buf" ]; then \
		cp -R generated/stubs/ts/buf ../app/frontend/src/gen/; \
	fi
	# Fix import paths in generated TypeScript files (buf/validate paths need adjustment after copying)
	@find ../app/frontend/src/gen -name "*.ts" -type f -exec sed -i 's|../../../../buf/validate/validate_pb|../../buf/validate/validate_pb|g' {} \;
	# Copy Connect client directories (they're in subdirectories like cloudresourcev1connect/)
	@for dir in generated/stubs/go/github.com/project-planton/project-planton/apis/org/project_planton/app/*/v1/*connect; do \
		if [ -d "$$dir" ]; then \
			service=$$(echo "$$dir" | sed 's|.*/app/\([^/]*\)/v1/.*|\1|'); \
			connectdir=$$(basename "$$dir"); \
			dest="org/project_planton/app/$$service/v1/$$connectdir"; \
			mkdir -p "$$dest"; \
			cp -r "$$dir"/* "$$dest/"; \
		fi; \
	done
	# Copy backend Connect clients and proto files to app/backend/apis/gen/go/proto
	@if [ -d "generated/stubs/go/github.com/project-planton/project-planton/apis/org/project_planton/app" ]; then \
		mkdir -p ../app/backend/apis/gen/go/proto/backendv1connect; \
		rm -rf ../app/backend/apis/gen/go/proto/*.go; \
		rm -rf ../app/backend/apis/gen/go/proto/backendv1connect/*.go; \
		for dir in generated/stubs/go/github.com/project-planton/project-planton/apis/org/project_planton/app/*/v1; do \
			if [ -d "$$dir" ]; then \
				find "$$dir" -name "*connect.go" -exec cp {} ../app/backend/apis/gen/go/proto/backendv1connect/ \; 2>/dev/null || true; \
				find "$$dir" -name "*.pb.go" -exec cp {} ../app/backend/apis/gen/go/proto/ \; 2>/dev/null || true; \
			fi; \
		done; \
	fi
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
