package service

import (
	"fmt"
	"os"

	ilog "github.com/jinmukeji/plat-pkg/v4/rpc/internal/log"
	"github.com/jinmukeji/plat-pkg/v4/rpc/internal/version"

	wcid "github.com/jinmukeji/plat-pkg/v4/micro/handler/cid"
	wlog "github.com/jinmukeji/plat-pkg/v4/micro/handler/log"
	wme "github.com/jinmukeji/plat-pkg/v4/micro/handler/microerr"

	wsvc "github.com/go-micro/plugins/v4/wrapper/service"
	"github.com/jinmukeji/plat-pkg/v4/rpc/internal/config"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"
)

type ServiceOptions struct {
	options

	// PreServerHandlerWrappers 自定义 HandlerWrapper，在标准 HandlerWrapper 之前注册
	PreServerHandlerWrappers []server.HandlerWrapper

	// PostServerHandlerWrappers 自定义 HandlerWrapper，在标准 HandlerWrapper 之后注册
	PostServerHandlerWrappers []server.HandlerWrapper

	// PreClientWrappers 自定义 Client Wrapper，在标准 Wrapper 之前注册
	PreClientWrappers []client.Wrapper

	// PostClientWrappers 自定义 Client Wrapper，在标准 Wrapper 之前注册
	PostClientWrappers []client.Wrapper

	// ServiceOptions 其它 Service Option
	ServiceOptions []micro.Option
}

func NewServiceOptions(namespace, name string) *ServiceOptions {
	o := ServiceOptions{}
	o.Namespace = namespace
	o.Name = name

	return &o
}

func CreateService(opts *ServiceOptions) micro.Service {
	// 设置 service，并且加载配置信息
	svc := newService(opts)
	err := setupService(svc, opts)
	die(err)

	return svc
}

func newService(opts *ServiceOptions) micro.Service {
	versionMeta := opts.ServiceMetadata()

	// Create a new service. Optionally include some options here.
	svcOpts := []micro.Option{
		// Service Basic Info
		micro.Name(opts.FQDN()),
		micro.Version(opts.ProductVersion),

		// Fault Tolerance - Heartbeating
		micro.RegisterTTL(defaultRegisterTTL),
		micro.RegisterInterval(defaultRegisterInterval),

		// Setup metadata
		micro.Metadata(versionMeta),
	}
	if len(opts.ServiceOptions) > 0 {
		svcOpts = append(svcOpts, opts.ServiceOptions...)
	}

	svc := micro.NewService(svcOpts...)

	svc.Options().Cmd.App().Description = fmt.Sprintf("fqdn: %s", opts.FQDN())

	return svc
}

func setupService(svc micro.Service, opts *ServiceOptions) error {
	// 设置启动参数
	flags := defaultFlags()
	if len(opts.Flags) > 0 {
		flags = append(flags, opts.Flags...)
	}

	svc.Init(
		micro.Flags(flags...),

		micro.Action(func(c *cli.Context) error {
			if opts.CliPreAction != nil {
				opts.CliPreAction(c)
			}

			if c.Bool("version") {
				version.PrintFullVersionInfo(opts)
				os.Exit(0)
			}

			ilog.SetupLogger(c, opts.Name)

			// 启动阶段打印版本号
			// 由于内部使用了 logger，需要在 logger 被设置后调用
			version.LogVersionInfo(opts)

			// 加载 config
			err := config.SetupConfig(c)
			if err != nil {
				return err
			}

			if opts.CliPostAction != nil {
				opts.CliPostAction(c)
			}

			return nil
		}),
	)

	// Setup wrappers
	setupHandlerWrappers(svc, opts)

	return nil
}

func defaultFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "version",
			Usage: "Show version information",
		},
	}

	flags = append(flags, ilog.MicroCliFlags()...)
	flags = append(flags, config.MicroCliFlags()...)

	return flags
}

func setupHandlerWrappers(svc micro.Service, opts *ServiceOptions) {
	// 设置 Server Handler Wrappers
	srvWrappers := []server.HandlerWrapper{}

	// 自定义 pre
	if len(opts.PreServerHandlerWrappers) > 0 {
		srvWrappers = append(srvWrappers, opts.PreServerHandlerWrappers...)
	}

	srvWrappers = append(srvWrappers,
		// 默认的 wrappers
		wsvc.NewHandlerWrapper(svc),
		wcid.CidWrapper,
		wme.MicroErrWrapper,
		wlog.LogWrapper,
	)

	// 自定义 post
	if len(opts.PostServerHandlerWrappers) > 0 {
		srvWrappers = append(srvWrappers, opts.PostServerHandlerWrappers...)
	}

	svc.Init(micro.WrapHandler(srvWrappers...))

	// 设置 Client Wrappers
	clientWrappers := []client.Wrapper{}
	if len(opts.PreClientWrappers) > 0 {
		clientWrappers = append(clientWrappers, opts.PreClientWrappers...)
	}

	clientWrappers = append(clientWrappers,
		// 默认的 wrappers
		wsvc.NewClientWrapper(svc),
	)
	if len(opts.PostClientWrappers) > 0 {
		clientWrappers = append(clientWrappers, opts.PostClientWrappers...)
	}

	svc.Init(
		micro.WrapClient(clientWrappers...),
	)
}
