package service

import (
	"fmt"
	"os"

	ilog "github.com/jinmukeji/plat-pkg/v4/rpc/internal/log"

	"github.com/jinmukeji/plat-pkg/v4/rpc/internal/config"
	"github.com/jinmukeji/plat-pkg/v4/rpc/internal/version"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/web"
)

type WebOptions struct {
	options

	// WebOptions 其它 Web Option
	WebOptions []web.Option
}

func NewWebOptions(namespace, name string) *WebOptions {
	o := WebOptions{}
	o.Namespace = namespace
	o.Name = name

	return &o
}

func CreateWeb(opts *WebOptions) web.Service {
	// 设置 service，并且加载配置信息
	svc := newWebService(opts)
	err := setupWebService(svc, opts)
	die(err)

	return svc
}

func newWebService(opts *WebOptions) web.Service {
	versionMeta := opts.ServiceMetadata()

	// Create a new service. Optionally include some options here.
	svcOpts := []web.Option{
		// Service Basic Info
		web.Name(opts.FQDN()),
		web.Version(opts.ProductVersion),

		// Fault Tolerance - Heartbeating
		web.RegisterTTL(defaultRegisterTTL),
		web.RegisterInterval(defaultRegisterInterval),

		// Setup metadata
		web.Metadata(versionMeta),
	}

	if len(opts.WebOptions) > 0 {
		svcOpts = append(svcOpts, opts.WebOptions...)
	}

	svc := web.NewService(svcOpts...)
	svc.Options().Service.Options().Cmd.App().Description = fmt.Sprintf("fqdn: %s", opts.FQDN())

	return svc
}

func setupWebService(svc web.Service, opts *WebOptions) error {
	// 设置启动参数
	flags := defaultWebFlags()
	if len(opts.Flags) > 0 {
		flags = append(flags, opts.Flags...)
	}

	err := svc.Init(
		// Setup runtime flags
		web.Flags(flags...),

		web.Action(func(c *cli.Context) {
			if opts.CliPreAction != nil {
				opts.CliPreAction(c)
			}

			if c.Bool("version") {
				version.PrintFullVersionInfo(opts)
				os.Exit(0)
			}

			ilog.SetupLogger(c, opts.Name)

			// 启动阶段打印版本号
			// 由于内部使用到了 logger，需要在 logger 被设置后调用
			version.LogVersionInfo(opts)

			// 加载 config
			err := config.SetupConfig(c)
			die(err)

			if opts.CliPostAction != nil {
				opts.CliPostAction(c)
			}
		}),
	)

	if err != nil {
		return err
	}

	return nil
}

func defaultWebFlags() []cli.Flag {
	return defaultFlags()
}
