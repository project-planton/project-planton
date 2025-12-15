version?=dev
name=project-planton
gcs_bucket=project-planton-downloads
name_local=project-planton
pkg=github.com/project-planton/project-planton
build_dir=build
LDFLAGS=-ldflags "-X ${pkg}/internal/cli/version.Version=${version}"
BAZEL?=./bazelw

# If BUILDBUDDY_API_KEY is set, enable the :bb config and inject only the header.
ifneq ($(strip $(BUILDBUDDY_API_KEY)),)
BAZEL_REMOTE_FLAGS=--config=bb --remote_header=x-buildbuddy-api-key=$$BUILDBUDDY_API_KEY
else
BAZEL_REMOTE_FLAGS=
endif

build_cmd=go build -v ${LDFLAGS}

PARALLEL?=$(shell getconf _NPROCESSORS_ONLN 2>/dev/null || sysctl -n hw.ncpu)

clean-bazel:
	rm -rf .bazelbsp bazel-bin bazel-out bazel-testlogs bazel-project-planton

reset-ide: clean-bazel
	rm -rf .idea

.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: build_darwin
build_darwin:
	GOOS=darwin ${build_cmd} -o ${build_dir}/${name}-darwin .

.PHONY: buf-generate
buf-generate: protos

.PHONY: protos
protos:
	pushd apis;make build;popd
	${BAZEL} run //:gazelle

.PHONY: buf-lint
buf-lint:
	$(MAKE) -C apis buf-lint

.PHONY: bazel-mod-tidy
bazel-mod-tidy:
	${BAZEL} mod tidy

.PHONY: gazelle
gazelle: bazel-gazelle

.PHONY: bazel-gazelle
bazel-gazelle:
	${BAZEL} run ${BAZEL_REMOTE_FLAGS} //:gazelle

.PHONY: clean-gazelle
clean-gazelle:
	@echo "Cleaning all BUILD.bazel files (excluding root)..."
	@find . -mindepth 2 -name "BUILD.bazel" -type f -delete
	@echo "✅ All BUILD.bazel files removed (root preserved)."

.PHONY: reset-gazelle
reset-gazelle: clean-gazelle gazelle
	@echo "✅ Gazelle reset complete. BUILD.bazel files regenerated."

.PHONY: bazel-build-cli
bazel-build-cli:
	${BAZEL} build ${BAZEL_REMOTE_FLAGS} //:project-planton

.PHONY: bazel-test
bazel-test:
	${BAZEL} test ${BAZEL_REMOTE_FLAGS} --test_output=errors //...

# Generates kind_map_gen.go containing ToMessageMap.
# The "-tags codegen" flag is REQUIRED to avoid chicken-and-egg compilation errors.
# See pkg/crkreflect/new_instance.go and pkg/crkreflect/codegen/main.go for details.
.PHONY: generate-cloud-resource-kind-map
generate-cloud-resource-kind-map:
	rm -f pkg/crkreflect/kind_map_gen.go
	go run -tags codegen ./pkg/crkreflect/codegen

.PHONY: generate-kubernetes-types
generate-kubernetes-types:
	pushd pkg/kubernetes/kubernetestypes;make build;popd

.PHONY: build-go
build-go: fmt deps vet
	GOOS=darwin GOARCH=amd64 ${build_cmd} -o ${build_dir}/${name}-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 ${build_cmd} -o ${build_dir}/${name}-darwin-arm64 .
	GOOS=linux GOARCH=amd64 ${build_cmd} -o ${build_dir}/${name}-linux .
	openssl dgst -sha256 ${build_dir}/${name}-darwin-arm64
	openssl dgst -sha256 ${build_dir}/${name}-linux

.PHONY: build-cli
build-cli: build-go

.PHONY: build
build: protos generate-cloud-resource-kind-map bazel-mod-tidy bazel-gazelle bazel-build-cli build-cli

${build_dir}/${name}: build-go

.PHONY: test
test:
	go test -race -v -count=1 -p $(PARALLEL) ./...

.PHONY: run
run: build
	${build_dir}/${name}

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf ${build_dir}

.PHONY: checksum_darwin
checksum_darwin:
	@openssl dgst -sha256 ${build_dir}/${name}-darwin

.PHONY: checksum_linux
checksum_linux:
	openssl dgst -sha256 ${build_dir}/${name}-linux

.PHONY: checksum
checksum: checksum_darwin checksum_linux

.PHONY: local
local: build_darwin
	rm -f ${HOME}/bin/${name_local}
	cp ./${build_dir}/${name}-darwin ${HOME}/bin/${name_local}
	chmod +x ${HOME}/bin/${name_local}

.PHONY: show-todo
show-todo:
	grep -r "TODO:" cmd internal

.PHONY: release-buf
release-buf:
	pushd apis;buf push;buf push --label ${version};popd

.PHONY: release-github
release-github:
	git tag ${version}
	git push origin ${version}
	@if ! command -v gh >/dev/null 2>&1; then \
		echo "GitHub CLI (gh) not found. Install from https://cli.github.com/"; \
		exit 1; \
	fi
	@if ! gh auth status >/dev/null 2>&1; then \
		echo "GitHub CLI is not authenticated. Run: gh auth login"; \
		exit 1; \
	fi
	@assets=""; \
	for f in ${build_dir}/${name}-darwin-amd64 ${build_dir}/${name}-darwin-arm64 ${build_dir}/${name}-linux apis/internal/generated/docs/docs.json; do \
		if [ -f "$$f" ]; then assets="$$assets $$f"; fi; \
	done; \
	echo "Creating GitHub release ${version} with assets: $$assets"; \
#	gh release create "${version}" --title "${version}" --generate-notes $$assets
	gh release create "${version}" --title "${version}" --generate-notes

.PHONY: release
release: protos build-cli test release-github release-buf

.PHONY: run-docs
run-docs:
	pushd docs;make run;popd

.PHONY: build-docs
build-docs:
	pushd docs;make build;popd

# ── website (site/) ────────────────────────────────────────────────────────────
.PHONY: run-site
run-site:
	$(MAKE) -C site dev

.PHONY: build-site
build-site:
	$(MAKE) -C site build

.PHONY: preview-site
preview-site:
	$(MAKE) -C site preview-site
