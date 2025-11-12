SHELL := /bin/bash

.PHONY: deps dev build start run lint clean preview-site export-static

PORT ?= 3000

deps:
	yarn install

dev:
	yarn dev -- --port $(PORT)

run: dev

copy-docs:
	yarn copy-docs

build: deps copy-docs
	yarn build

start:
	yarn start

lint:
	yarn lint

export-static: build

preview-site: export-static
	yarn serve -s out -l $(PORT)

clean:
	rm -rf .next out
