package jwt

import (
	"time"

	"github.com/jinmukeji/plat-pkg/v4/micro/meta"
)

type Options struct {
	Enabled         bool          // 是否启用
	MaxExpInterval  time.Duration // 最大过期时间间隔
	HeaderKey       string        // HTTP Request Header 中的 jwt 使用的 key
	MicroConfigPath string        // Micro Config 中的 key
}

const (
	DefaultMaxExpInterval  = 10 * time.Minute
	DefaultMicroConfigPath = "platform/app-key"
)

func DefaultOptions() Options {
	return Options{
		MaxExpInterval:  DefaultMaxExpInterval,
		HeaderKey:       meta.MetaKeyJwt,
		MicroConfigPath: DefaultMicroConfigPath,
	}
}
