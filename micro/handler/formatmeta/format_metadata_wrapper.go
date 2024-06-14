package formatmeta

import (
	"context"

	pm "github.com/jinmukeji/plat-pkg/v4/micro/meta"

	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
)

// FormatMetadataWrapper 格式化所有 metadata keys
func FormatMetadataWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		md, ok := metadata.FromContext(ctx)
		if ok && len(md) > 0 {
			nmd := metadata.Metadata{}
			for k, v := range md {
				nmd[pm.StandardizeKey(k)] = v
			}

			ctx = metadata.NewContext(ctx, nmd)
		}

		err := fn(ctx, req, rsp)
		return err
	}
}
