package mysql

import (
	"time"

	"github.com/go-sql-driver/mysql"

	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
)

type microDbOptions struct {
	// 网络方式, tcp 或者 udp
	Network string `json:"network" yaml:"network"`

	// 服务器地址 - host:port
	Address string `json:"address" yaml:"address"`

	// 用户名
	Username string `json:"username" yaml:"username"`

	// 密码
	Password string `json:"password" yaml:"password"`

	// 数据库名
	Database string `json:"database" yaml:"database"`

	// 连接超时
	// The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
	DialTimeout string `json:"dial_timeout" yaml:"dial_timeout"`

	// 是否转换时间
	ParseTime *bool `json:"parse_time,omitempty" yaml:"parse_time,omitempty"`

	// 字符排序
	Collation string `json:"collation" yaml:"collation"`

	// 区域设置
	Locale string `json:"locale" yaml:"locale"`
}

// NewDbClientFromConfig 通过 Micro Config 的配置创建 DbClient
func NewConfigFromMicroConfigKey(cfgKey ...string) (*mysql.Config, error) {
	opts := microDbOptions{}
	if err := config.Get(cfgKey...).Scan(&opts); err != nil {
		return nil, err
	}

	return mapConfig(&opts)
}

// NewConfigFromMicroConfigValue 通过 Micro Config 的 reader.Value 创建 DbClient
func NewConfigFromMicroConfigValue(v reader.Value) (*mysql.Config, error) {
	opts := microDbOptions{}
	if err := v.Scan(&opts); err != nil {
		return nil, err
	}

	return mapConfig(&opts)
}

func mapConfig(opts *microDbOptions) (*mysql.Config, error) {
	cfg := NewConfig()

	if len(opts.Network) > 0 {
		cfg.Net = opts.Network
	}
	cfg.Addr = opts.Address
	cfg.User = opts.Username
	cfg.Passwd = opts.Password
	cfg.DBName = opts.Database

	if len(opts.DialTimeout) > 0 {
		timeout, err := time.ParseDuration(opts.DialTimeout)
		if err != nil {
			return nil, err
		}
		cfg.Timeout = timeout
	}

	if opts.ParseTime != nil {
		cfg.ParseTime = *opts.ParseTime
	}

	if len(opts.Collation) > 0 {
		cfg.Collation = opts.Collation
	}

	if len(opts.Locale) > 0 {
		loc, err := time.LoadLocation(opts.Locale)
		if err != nil {
			return nil, err
		}
		cfg.Loc = loc
	}

	return cfg, nil
}
