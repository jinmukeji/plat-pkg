module github.com/jinmukeji/plat-pkg/v2

go 1.12

replace (
	// FIXME: 由于 etcd 与 gRPC 的兼容问题，暂时使用定制的 etcd 版本
	// https://github.com/etcd-io/etcd/issues/11721
	github.com/coreos/etcd => github.com/skyjia/etcd v3.3.22-grpc1.27-origmodule+incompatible
	github.com/micro/cli/v2 => github.com/jinmukeji/jm-micro/v2 v2.0.0

	github.com/micro/go-micro/v2 => github.com/micro/go-micro/v2 v2.8.0
	github.com/micro/micro/v2 => github.com/micro/micro/v2 v2.8.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

