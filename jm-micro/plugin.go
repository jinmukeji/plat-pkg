package main

import (
	"github.com/jinmukeji/plat-pkg/v2/jm-micro/plugins/cid"
	"github.com/jinmukeji/plat-pkg/v2/jm-micro/plugins/configloader"
	"github.com/jinmukeji/plat-pkg/v2/jm-micro/plugins/healthcheck"
	"github.com/jinmukeji/plat-pkg/v2/jm-micro/plugins/jwt"
	"github.com/jinmukeji/plat-pkg/v2/jm-micro/plugins/log"
	"github.com/jinmukeji/plat-pkg/v2/jm-micro/plugins/tcphealthcheck"

	// NOTE: go-micro/v2 标准框架已经默认采用 gRPC，TLS 设定方式发生变化，本插件不适用于 gRPC.
	// "github.com/jinmukeji/plat-pkg/v2/jm-micro/plugins/tls-client"

	// "github.com/jinmukeji/plat-pkg/v2/jm-micro/plugins/whitelist"

	"github.com/go-micro/plugins/v2/micro/cors"
	// "github.com/micro/go-plugins/micro/gzip"
	"github.com/go-micro/plugins/v2/micro/metadata"

	"github.com/micro/micro/v2/client/api"
	"github.com/micro/micro/v2/plugin"
)

func init() {
	// 初始化 micro 插件

	// 全局插件，针对全部子命令
	initGlobalPlugins()

	// proxy 服务插件
	initProxyPlugins()

	// web 服务插件
	initWebPlugins()

	// api 服务插件
	initAPIPlugins()
}

func initGlobalPlugins() {
	// 全局插件
	err := plugin.Register(log.NewPlugin(Name))
	die(err)

	err = plugin.Register(configloader.NewPlugin())
	die(err)

	err = plugin.Register(metadata.NewPlugin())
	die(err)

	err = plugin.Register(cors.NewPlugin())
	die(err)
}

func initProxyPlugins() {
	// proxy 服务插件
	// err := proxy.Register(tls.NewPlugin())
	// die(err)
}

func initWebPlugins() {
	// web 服务插件
	// err := web.Register(tls.NewPlugin())
	// die(err)
}

func initAPIPlugins() {
	// api 服务插件

	// err = api.Register(tls.NewPlugin())
	// die(err)

	// micro gzip 插件存在 bug，当 response 数据量过小的时候，压缩后的数据丢失
	// err = api.Register(gzip.NewPlugin())
	// die(err)
	err := api.Register(healthcheck.NewPlugin())
	die(err)

	err = api.Register(tcphealthcheck.NewPlugin())
	die(err)

	err = api.Register(cid.NewPlugin())
	die(err)

	err = api.Register(jwt.NewPlugin())
	die(err)

	// TODO: 白名单插件
	// err = api.Register(whitelist.NewRPCWhitelist("com.jinmuhealth.platform.srv.template-service1"))
	// die(err)
}
