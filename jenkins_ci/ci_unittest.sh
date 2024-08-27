#!/bin/bash
#单元测试
CUR=$(dirname $0)
cd ${CUR}/..
ETCD_ADDR=${DOCKER_HOST_IP}:2379 make test
