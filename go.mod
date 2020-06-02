module github.com/jinmukeji/plat-pkg/v2

go 1.14

replace (
	// FIXME: 由于 etcd 与 gRPC 的兼容问题，暂时使用定制的 etcd 版本
	// https://github.com/etcd-io/etcd/issues/11721
	github.com/coreos/etcd => github.com/skyjia/etcd v3.3.22-grpc1.27-origmodule+incompatible

	github.com/micro/go-micro/v2 => github.com/micro/go-micro/v2 v2.8.0
	github.com/micro/micro/v2 => github.com/micro/micro/v2 v2.8.0
)

require (
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gobwas/glob v0.2.3
	github.com/jinmukeji/go-pkg/v2 v2.4.1
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.8.0
	github.com/micro/go-plugins/logger/logrus/v2 v2.5.0
	github.com/micro/go-plugins/micro/cors/v2 v2.8.0
	github.com/micro/go-plugins/micro/metadata/v2 v2.5.0
	github.com/micro/go-plugins/wrapper/service/v2 v2.5.0
	github.com/micro/micro/v2 v2.8.0
	github.com/rs/xid v1.2.1
	github.com/sirupsen/logrus v1.6.0
	github.com/smallstep/cli v0.14.4
	github.com/stretchr/testify v1.6.0
	gopkg.in/yaml.v2 v2.3.0
)
