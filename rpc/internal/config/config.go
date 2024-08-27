package config

import (
	"fmt"
	"sync"

	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"github.com/go-micro/plugins/v4/config/source/etcd"
	mlog "github.com/jinmukeji/go-pkg/v2/log"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
)

var (
	log = mlog.StandardLogger()

	yamlConfig config.Config

	yamlEncoder = yaml.NewEncoder()

	lock sync.Once
)

// Config 相关常量
const (
	DefaultConfigEtcdPrefix = "/micro/config/jm/"
)

// YamlConfig 获取 yaml 配置解析器
func YamlConfig() config.Config {
	lock.Do(func() {
		// 创建 yaml 解析器
		yamlConfig, _ = config.NewConfig(
			config.WithReader(json.NewReader(reader.WithEncoder(yamlEncoder))),
		)
	})

	return yamlConfig
}

func MicroCliFlags() []cli.Flag {

	return []cli.Flag{
		// Config 相关
		&cli.StringSliceFlag{
			Name:  "config_file",
			Usage: "Config file path",
		},

		&cli.StringFlag{
			Name:  "config_etcd_address",
			Usage: "Etcd config source address",
		},

		&cli.StringFlag{
			Name:  "config_etcd_prefix",
			Usage: "Etcd config K/V prefix",
			Value: DefaultConfigEtcdPrefix,
		},
	}
}

func SetupConfig(c *cli.Context) error {
	// 加载以下配置信息数据源，优先级依次从低到高：
	// 1. Etcd K/V 配置中心
	// 2. 配置文件，YAML格式
	// 3. 环境变量（暂不实现）

	cfgEtcdAddr := c.String("config_etcd_address")
	cfgEtcdPrefix := c.String("config_etcd_prefix")

	// Load config from etcd
	if cfgEtcdAddr != "" {
		etcdSource := etcd.NewSource(
			etcd.WithAddress(cfgEtcdAddr),
			etcd.WithPrefix(cfgEtcdPrefix),
			source.WithEncoder(yamlEncoder),
		)

		if err := YamlConfig().Load(etcdSource); err != nil {
			return fmt.Errorf("failed to load config from etcd at %s with prefix of [%s]: %w", cfgEtcdAddr, cfgEtcdPrefix, err)
		}

		log.Infof("Loaded config from etcd at %s with prefix of [%s]", cfgEtcdAddr, cfgEtcdPrefix)
	}

	// Load config from files
	cfgFiles := c.StringSlice("config_file")
	for _, f := range cfgFiles {
		fileSource := file.NewSource(
			file.WithPath(f),
			source.WithEncoder(yamlEncoder),
		)

		if err := YamlConfig().Load(fileSource); err != nil {
			return fmt.Errorf("failed to load config file %s: %w", f, err)
		}

		log.Infof("Loaded config from file: %s", f)
	}

	return nil
}
