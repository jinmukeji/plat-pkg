module github.com/jinmukeji/plat-pkg/v2

go 1.16

replace (
	// FIXME: 由于 etcd 与 gRPC 的兼容问题，暂时使用定制的 etcd 版本
	//  https://github.com/etcd-io/etcd/issues/11721
	//  https://github.com/etcd-io/etcd/issues/11154
	//  https://github.com/etcd-io/etcd/pull/11823
	github.com/coreos/etcd => github.com/skyjia/etcd v3.3.22-grpc1.27-origmodule+incompatible

	// FIXME: etcd 与 grpc 1.30+ 版本不兼容，此处降级到 v0.15.2
	github.com/smallstep/cli => github.com/smallstep/cli v0.15.2
	// FIXME: etcd 与 grpc 1.30+ 版本不兼容，此处降级到 1.29
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)

require (
	github.com/coreos/etcd v3.3.27+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.9.0
	github.com/go-micro/plugins/v2/logger/logrus v0.0.0-20220711144004-e3081cf21c80
	github.com/go-micro/plugins/v2/micro/cors v0.0.0-20220711144004-e3081cf21c80
	github.com/go-micro/plugins/v2/micro/metadata v0.0.0-20220711144004-e3081cf21c80
	github.com/go-micro/plugins/v2/wrapper/service v0.0.0-20220711144004-e3081cf21c80
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gobwas/glob v0.2.3
	github.com/jinmukeji/go-pkg/v2 v2.6.0
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.3
	github.com/rs/xid v1.4.0
	github.com/sirupsen/logrus v1.8.1
	github.com/smallstep/cli v0.15.0
	github.com/stretchr/testify v1.8.1
	github.com/ugorji/go v1.2.7 // indirect
	google.golang.org/grpc v1.44.0
	gopkg.in/yaml.v2 v2.4.0
)
