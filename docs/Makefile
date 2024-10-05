.PHONY: clean
clean:
	rm -rf node_modules .next .yarn

.PHONY: deps
deps:
	yarn install

.PHONY: build
build: clean deps
	yarn build

.PHONY: run
run: clean deps
	yarn dev
