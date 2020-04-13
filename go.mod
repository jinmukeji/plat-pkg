module github.com/jinmukeji/plat-pkg/v2

go 1.14

replace (
	// FIXME: 修复 go-micro v2.4.0 相关包引用的 bug
	github.com/micro/go-micro/v2 => github.com/micro/go-micro/v2 v2.4.0
	github.com/micro/micro/v2 => github.com/micro/micro/v2 v2.4.0

	// FIXME: smallstep/cli v0.14.2 的 GOSUM 校验失败
	github.com/smallstep/cli => github.com/smallstep/cli v0.13.3

	// FIXME: 由于 etcd 与 gRPC 的兼容问题，得降级 grpc 版本
	// https://github.com/etcd-io/etcd/issues/11721
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gobwas/glob v0.2.3
	github.com/jinmukeji/go-pkg/v2 v2.2.7
	github.com/jinzhu/gorm v1.9.12
	github.com/manifoldco/promptui v0.7.0 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.4.0
	github.com/micro/go-plugins/logger/logrus/v2 v2.3.0
	github.com/micro/go-plugins/micro/cors/v2 v2.3.0
	github.com/micro/go-plugins/micro/metadata/v2 v2.3.0
	github.com/micro/go-plugins/wrapper/service/v2 v2.3.0
	github.com/micro/micro/v2 v2.4.0
	github.com/samfoo/ansi v0.0.0-20160124022901-b6bd2ded7189 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/smallstep/assert v0.0.0-20200103212524-b99dc1097b15 // indirect
	github.com/smallstep/cli v0.13.3
	github.com/stretchr/testify v1.5.1
	go.etcd.io/etcd v3.3.20+incompatible
	gopkg.in/yaml.v2 v2.2.8
)
