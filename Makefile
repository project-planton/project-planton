version?=dev
name=project-planton
gcs_bucket=project-planton-downloads
name_local=project-planton
pkg=github.com/project-planton/project-planton
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

.PHONY: protos
protos:
	pushd apis;make build;popd

.PHONY: generate-kubernetes-types
generate-kubernetes-types:
	pushd pkg/kubernetestypes;make build;popd

.PHONY: build-cli
build-cli: ${build_dir}/${name}

.PHONY: build
build: protos build-cli

${build_dir}/${name}: deps vet
	GOOS=darwin GOARCH=amd64 ${build_cmd} -o ${build_dir}/${name}-darwin-amd64 .
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

.PHONY: upload-cli-binaries
upload-cli-binaries:
	gsutil -h "Cache-Control:no-cache" cp build/${name}-linux gs://${gcs_bucket}/cli/${version}/${name}-${version}-linux
	gsutil -h "Cache-Control:no-cache" cp build/${name}-darwin-amd64 gs://${gcs_bucket}/cli/${version}/${name}-${version}-amd64
	gsutil -h "Cache-Control:no-cache" cp build/${name}-darwin-arm64 gs://${gcs_bucket}/cli/${version}/${name}-${version}-arm64

.PHONY: release-buf
release-buf:
	pushd apis;buf push;buf push --label ${version};popd

.PHONY: release-github
release-github:
	git tag ${version}
	git push origin ${version}
	gh release create ${version} \
         --title ${version} \
         --target ${version} \
         build/project-planton-darwin-amd64 \
         build/project-planton-darwin-arm64 \
         build/project-planton-linux \
         apis/internal/generated/docs/docs.json

.PHONY: release
release: protos release-buf build-cli upload-cli-binaries release-github

.PHONY: run-docs
run-docs:
	pushd docs;make run;popd

.PHONY: build-docs
build-docs:
	pushd docs;make build;popd
