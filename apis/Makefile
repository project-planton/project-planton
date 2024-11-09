.PHONY: build
build:
	rm -rf internal/generated go
	mkdir -p internal/generated go
	buf generate
	cp -R internal/generated/go/github.com/project-planton/project-planton/apis/go/. go/
	pushd go;go vet ./...;go fmt ./...;popd
