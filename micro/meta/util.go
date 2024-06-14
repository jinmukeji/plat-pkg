package meta

import (
	"context"
	"strings"

	"go-micro.dev/v4/metadata"
)

// StandardizeKey 标准化 metadata 中使用的 key
func StandardizeKey(key string) string {
	return strings.ToLower(key)
}

func Get(ctx context.Context, key string) (string, bool) {
	return metadata.Get(ctx, key)
}

func MustGet(ctx context.Context, key string) string {
	if v, ok := metadata.Get(ctx, key); ok {
		return v
	}
	return ""
}

func Set(ctx context.Context, key, value string) context.Context {
	return metadata.Set(ctx, key, value)
}

func Delete(ctx context.Context, key string) context.Context {
	return metadata.Delete(ctx, key)
}
