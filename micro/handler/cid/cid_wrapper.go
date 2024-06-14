package cid

import (
	"context"

	"github.com/jinmukeji/plat-pkg/v4/micro/meta"
	"github.com/jinmukeji/plat-pkg/v4/micro/tracer"
	"go-micro.dev/v4/server"
)

// CidWrapper 如果请求中没有 cid，则生成一个新的
func CidWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		cid := meta.CidFromContext(ctx)
		// 如果没有找到 cid，则生成一个新的
		if cid == "" {
			cid = tracer.NewCid()
			ctx = meta.ContextWithCid(ctx, cid)
		}
		err := fn(ctx, req, rsp)
		return err
	}
}
