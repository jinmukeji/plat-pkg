package mysql

import (
	"time"

	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB = gorm.DB

const (
	tlsKey = "custom"
)

// OpenDB is alias of OpenGormDB.
func OpenDB(opt ...Option) (*DB, error) {
	return OpenGormDB(opt...)
}

// OpenGormDB 打开一个 *gorm.DB 的连接
func OpenGormDB(opt ...Option) (*DB, error) {
	options := NewOptions(opt...)

	mysqlCfg := options.MySqlCfg

	if options.TLSCfg != nil {
		err := mysql.RegisterTLSConfig(tlsKey, options.TLSCfg)
		if err != nil {
			return nil, err
		}

		mysqlCfg.TLSConfig = tlsKey
	}

	dsn := mysqlCfg.FormatDSN()

	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// 单数形式命名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// 禁止没有 WHERE 语句的 DELETE 或 UPDATE 操作执行，否则抛出 error
		AllowGlobalUpdate: false,
		// 重置 SetNow 的时间获取方式为总是获取 UTC 时区时间
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
