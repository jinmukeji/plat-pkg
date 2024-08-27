#!/bin/bash
# 代码格式化
CUR=$(dirname $0)
cd ${CUR}/..
set -e
make format
git --no-pager diff HEAD
git --no-pager diff-index --quiet HEAD
