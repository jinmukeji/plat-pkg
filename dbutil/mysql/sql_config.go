package mysql

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

// NewConfig 创建一个标准的 *mysql.Config
func NewConfig() *mysql.Config {
	cfg := mysql.NewConfig()

	// 显式设定以下关键参数
	cfg.ParseTime = true
	cfg.Collation = "utf8mb4_general_ci"
	cfg.Loc = time.UTC
	cfg.Net = "tcp"

	return cfg
}
