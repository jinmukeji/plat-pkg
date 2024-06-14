#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`
CONFIG_FILE="${CUR}/error_codes.yml"
OUT_FILE="${CUR}/codes.gen.go"
PKG_NAME="codes"

cd ${CUR}

echo "Generating rpc/errors/codes.gen.go"
go run ./gen/gen.go -c ${CONFIG_FILE} -o ${OUT_FILE} -p ${PKG_NAME}
goimports -w ${OUT_FILE} 