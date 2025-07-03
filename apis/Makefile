.PHONY: build
build:
	rm -rf internal/generated
	mkdir -p internal/generated
	buf generate
	cp -R internal/generated/go/github.com/project-planton/project-planton/apis/. .
	rm -rf internal
