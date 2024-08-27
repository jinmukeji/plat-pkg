#!/bin/bash
#关闭服务
CUR=$(dirname $0)
cd ${CUR}/..
docker stop etcd || echo "etcd stop"
docker network rm default_network || echo "default_network stop"
