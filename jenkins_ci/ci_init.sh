#!/bin/bash
# 初始化项目
go version
# set -e
CUR=$(dirname $0)
cd ${CUR}/..
curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b $GOPATH/bin v1.59.1
golangci-lint --version

git config --global --add url."git@github.com:".insteadOf "https://github.com/"
go mod tidy
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/goreleaser/goreleaser/v2@latest
