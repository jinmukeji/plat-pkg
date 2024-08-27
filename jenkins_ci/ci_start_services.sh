#!/bin/bash
# 启动服务
pip3 install awscli --upgrade --user
export PATH=/root/.local/bin/:$PATH
aws --version
#docker login
$(aws ecr get-login --no-include-email)
#关闭已开启的服务
docker stop etcd || echo "etcd stop"
docker network rm default_network || echo "default_network stop"
CUR=$(dirname $0)
cd ${CUR}/
docker network create default_network
# 启动etcd
docker run \
    -d \
    --rm \
    -p 2379:2379 \
    -p 2380:2380 \
    --name etcd \
    --network=default_network \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd:2379 \
    bitnami/etcd:latest
cd ${CUR}/../jenkins_ci
sleep 2s
# 将etcd配置信息注入
wget https://s3.cn-north-1.amazonaws.com.cn/res.jinmuhealth.com/download/tools/etcddump/etcddump_0.1.3_Linux_x86_64.tar.gz
tar -zxvf etcddump_0.1.3_Linux_x86_64.tar.gz
# 在本地Deployment-config-kv下local ，执行put-etcd-all.sh ，通过
#     etcddump dump \
# 	--address=127.0.0.1:2379 \
# 	--prefix="/micro/config/jm" \
# 	--output=test.out
# #这个命令将配置dump到test.out文件
#通过restore导入配置信息
./etcddump restore \
    --address=${DOCKER_HOST_IP}:2379 \
    --file=test.out
