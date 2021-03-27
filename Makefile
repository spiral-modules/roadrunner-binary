#!/usr/bin/make
# Makefile manual (ru): <http://linux.yaroslavl.ru/docs/prog/gnu_make_3-79_russian_manual.html>
# Makefile manual (en): <https://www.gnu.org/software/make/manual/html_node/index.html#SEC_Contents>

SHELL = /bin/sh
GO_VERSION = 1.16.2
DOCKER_GO_RUN_ARGS = --rm -v "$(shell pwd):/src" -w "/src" -u "$(shell id -u):$(shell id -g)" -e "GOPATH=/tmp" -e "HOME=/tmp"

.PHONY : help build fmt lint gotest test
.DEFAULT_GOAL : help

# This will output the help for each task. Thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[32m%-11s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build app binary file
	docker run $(DOCKER_GO_RUN_ARGS) -e "CGO_ENABLED=0" "golang:$(GO_VERSION)-buster" \
		go build -trimpath -ldflags "-s" -o ./rr ./cmd/rr

fmt: ## Run source code formatter tools
	docker run $(DOCKER_GO_RUN_ARGS) -e "GO111MODULE=off" "golang:$(GO_VERSION)-buster" \
		sh -c 'go get golang.org/x/tools/cmd/goimports && $$GOPATH/bin/goimports -d -w .'
	docker run $(DOCKER_GO_RUN_ARGS) "golang:$(GO_VERSION)-buster" \
		sh -c 'gofmt -s -w -d .; go mod tidy'

lint: ## Run app linters
	docker run $(DOCKER_GO_RUN_ARGS) -t golangci/golangci-lint:v1.39-alpine golangci-lint run

gotest: ## Run app tests
	docker run $(DOCKER_GO_RUN_ARGS) -t "golang:$(GO_VERSION)-buster" go test -v -race -timeout 10s ./...

test: lint gotest ## Run app tests and linters
