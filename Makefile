SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GOPATH := $(shell go env GOPATH)
export GO111MODULE := on
export GOPROXY := https://goproxy.io,direct
export GOPRIVATE := github.com/jinmukeji/*

# Install all the build and lint dependencies
setup:
	# golangci-lint
	brew install golangci-lint
	# misspell
	# curl -L https://git.io/misspell | sh
.PHONY: setup

# Update go packages
go-mod-update:
	@echo "Checking updated go packages..."
	@go list -u -m all
	@echo "Updating go packages..."
	@go get -u -t ./...
	@$(MAKE) go-mod-tidy
.PHONY: go-update

# Clean go.mod
go-mod-tidy:
	@go mod tidy -v
	# git --no-pager diff HEAD
	# git --no-pager diff-index --quiet HEAD
.PHONY: go-mod-tidy

# Reset go.mod
go-mod-reset:
	@rm -f go.sum
	@sed -i '' -e '/^require/,/^)/d' go.mod
	@go mod tidy -v
.PHONY: go-mod-tidy

generate:
	@go generate ./...
.PHONY: generate

# Format go files
format:
	@goimports -w ./
.PHONY: format

# Run all the linters
lint:
	@golangci-lint run
.PHONY: lint

# Go build all
build:
	@go build ./... > /dev/null
.PHONY: build

# Go test all
test:
	@go test ./...
.PHONY: test

# Run all code checks
ci: generate format lint build test
.PHONY: ci

.DEFAULT_GOAL := ci
