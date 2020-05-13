module github.com/jinmukeji/plat-pkg/v2

go 1.14

replace (
	// FIXME: 由于 etcd 与 gRPC 的兼容问题，暂时使用定制的 etcd 版本
	// https://github.com/etcd-io/etcd/issues/11721
	github.com/coreos/etcd => github.com/skyjia/etcd v3.3.20-grpc1.27-origmodule+incompatible

	github.com/micro/go-micro/v2 => github.com/micro/go-micro/v2 v2.6.0
	github.com/micro/micro/v2 => github.com/micro/micro/v2 v2.6.0
)

require (
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gobwas/glob v0.2.3
	github.com/jinmukeji/go-pkg/v2 v2.2.7
	github.com/jinzhu/gorm v1.9.12
	github.com/juju/ansiterm v0.0.0-20180109212912-720a0952cc2a // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.6.0
	github.com/micro/go-plugins/logger/logrus/v2 v2.5.0
	github.com/micro/go-plugins/micro/cors/v2 v2.5.0
	github.com/micro/go-plugins/micro/metadata/v2 v2.5.0
	github.com/micro/go-plugins/wrapper/service/v2 v2.5.0
	github.com/micro/micro/v2 v2.5.0
	github.com/sirupsen/logrus v1.6.0
	github.com/smallstep/cli v0.14.3
	github.com/stretchr/testify v1.5.1
	gopkg.in/yaml.v2 v2.2.8
)
