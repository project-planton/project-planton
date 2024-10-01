version?=v0.0.1
name=project-planton
name_local=project-planton
pkg=github.com/plantoncloud/project-planton
build_dir=build
LDFLAGS=-ldflags "-X ${pkg}/internal/cli/version.Version=${version}"

build_cmd=go build -v ${LDFLAGS}

.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: build_darwin
build_darwin: vet
	GOOS=darwin ${build_cmd} -o ${build_dir}/${name}-darwin .

.PHONY: build-apis
build-apis:
	buf generate

.PHONY: build-cli
build-cli: ${build_dir}/${name}

.PHONY: build
build: build-apis build-cli

${build_dir}/${name}: deps vet
	GOOS=darwin GOARCH=arm64 ${build_cmd} -o ${build_dir}/${name}-darwin-arm64 .
	GOOS=linux GOARCH=amd64 ${build_cmd} -o ${build_dir}/${name}-linux .
	openssl dgst -sha256 ${build_dir}/${name}-darwin-arm64
	openssl dgst -sha256 ${build_dir}/${name}-linux

.PHONY: test
test:
	go test -race -v -count=1 ./...

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

.PHONY: upload-binaries
upload-binaries:
	gsutil -h "Cache-Control:no-cache" cp build/${name}-linux gs://${name}/cli/${version}/${name}-${version}-linux
	gsutil -h "Cache-Control:no-cache" cp build/${name}-darwin-amd64 gs://${name}/cli/${version}/${name}-${version}-amd64
	gsutil -h "Cache-Control:no-cache" cp build/${name}-darwin-arm64 gs://${name}/cli/${version}/${name}-${version}-arm64

.PHONY: release-buf
release-buf:
	buf push
	buf push --label ${version}

.PHONY: release-github
release-github:
	git tag ${version}
	git push origin ${version}

.PHONY: release
release: build upload-binaries release-buf release-github

.PHONY: update-deps
update-deps:
	go get github.com/plantoncloud-inc/planton-cloud-apis@latest
