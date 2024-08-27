#!/bin/bash
# 静态代码分析
CUR=$(dirname $0)
cd ${CUR}/..
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
go mod vendor
export GO111MODULE=on

# make lint
