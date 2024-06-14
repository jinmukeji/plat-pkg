package mysql

import (
	"context"
	"fmt"

	"github.com/jinmukeji/plat-pkg/v4/dbutil/mysql"
	"github.com/jinmukeji/plat-pkg/v4/store"
)

var _ store.Store = (*MySqlStore)(nil)

type MySqlStore struct {
	*mysql.DB
}

func NewMySqlStore(db *mysql.DB) *MySqlStore {
	return &MySqlStore{
		DB: db,
	}
}

func NewStore(db *mysql.DB) *MySqlStore {
	return NewMySqlStore(db)
}

// 实现 store.Closer 接口
func (s *MySqlStore) Close() error {
	if s.DB != nil {
		sql, _ := s.DB.DB()
		return sql.Close()
	}
	return nil
}

// 实现事务控制接口 store.Tx

type txDbCtxKey struct{}

func (s *MySqlStore) GetDB(ctx context.Context) *mysql.DB {
	if v := ctx.Value(txDbCtxKey{}); v != nil {
		return v.(*mysql.DB)
	} else {
		return s.DB
	}
}

func (s *MySqlStore) BeginTx(ctx context.Context) context.Context {
	db := s.Begin() // 从原始的 DB 开启事务
	return context.WithValue(ctx, txDbCtxKey{}, db)
}

func (s *MySqlStore) CommitTx(ctx context.Context) {
	db := s.GetDB(ctx)
	db.Commit()
}

func (s *MySqlStore) RollbackTx(ctx context.Context) {
	db := s.GetDB(ctx)
	db.Rollback()
}

func (s *MySqlStore) GetError(ctx context.Context) error {
	db := s.GetDB(ctx)
	return db.Error
}

func (s *MySqlStore) Transaction(ctx context.Context, fc func(txs *MySqlStore) error) (err error) {
	tx := s.BeginTx(ctx)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
			s.RollbackTx(tx)
			return
		}
	}()

	err = fc(s)
	if err == nil {
		s.CommitTx(tx)
		err = s.GetError(tx)
	}

	if err != nil {
		s.RollbackTx(tx)
	}

	return
}
