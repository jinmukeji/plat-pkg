package rpc

import (
	mconfig "github.com/jinmukeji/plat-pkg/v4/rpc/internal/config"
	"go-micro.dev/v4/config"
)

// YamlConfig 获取内部 yaml 解析器
func YamlConfig() config.Config {
	return mconfig.YamlConfig()
}
