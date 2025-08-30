.PHONY: build
build:
	rm -rf generated/stubs generated/docs
	mkdir -p generated/stubs generated/docs
	buf generate
	cp -R generated/stubs/go/github.com/project-planton/project-planton/apis/. .
	rm -rf generated/stubs
