.PHONY: deps
deps:
	yarn install

.PHONY: generate
generate:
	buf generate

.PHONY: build
build: deps
	yarn build

.PHONY: dev
dev: deps
	yarn dev

.PHONY: clean
clean:
	rm -rf node_modules .next .yarn build

.PHONY: update-deps
update-deps:
	yarn install

