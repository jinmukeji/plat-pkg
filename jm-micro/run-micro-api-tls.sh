#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

SERVER_IP=0.0.0.0
SERVER_PORT=8080
ETCD_ADDR=localhost:2379

echo "TLS is not ready by now."
exit 1

go run . \
    --log_level=DEBUG \
    --register_interval=5 \
    --register_ttl=10 \
    --client_pool_size=100 \
    --server_name=com.jinmuhealth.platform.api \
    --metadata=X-Err-Style=MicroDetailed \
    --enable_tls \
    --tls_cert_file=./cert/server/localhost/localhost.crt \
    --tls_key_file=./cert/server/localhost/localhost.key \
    --tls_client_ca_file=./cert/root_ca.crt \
    --config_etcd_address=${ETCD_ADDR} \
    api \
    --address=${SERVER_IP}:${SERVER_PORT} \
    --handler=rpc \
    --enable_rpc \
    --namespace=com.jinmuhealth.platform.srv \
    --enable_tcp_healthcheck
    --tcp_healthcheck_addr=:9901
    # --enable_jwt \
    # --jwt_max_exp_interval=600s \
    # --rpc_whitelist=com.jinmuhealth.platform.srv.template-service
