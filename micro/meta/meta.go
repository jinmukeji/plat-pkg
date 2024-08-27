package meta

import "context"

const (
	// MetaKeyCid Metadata 中 cid 的 key
	MetaKeyCid = "x-cid"

	// MetaKeyJwt Metadata 中 jwt 的 key
	MetaKeyJwt = "x-jwt"

	// MetaKeyAppId Metadata 中 APP ID 的 key
	MetaKeyAppId = "x-app-id"
)

func ContextWithCid(ctx context.Context, cid string) context.Context {
	return Set(ctx, MetaKeyCid, cid)
}

// CidFromContext 从 Context 中获取 cid 的值
func CidFromContext(ctx context.Context) string {
	return MustGet(ctx, MetaKeyCid)
}

// ContextWithJwt 向 Context (Metadata) 中存入 jwt 的值
func ContextWithJwt(ctx context.Context, jwt string) context.Context {
	return Set(ctx, MetaKeyJwt, jwt)
}

// JwtFromContext 从 Context 中获取 jwt 的值
func JwtFromContext(ctx context.Context) string {
	return MustGet(ctx, MetaKeyJwt)
}

// ContextWithAppId 向 Context (Metadata) 中存入 APP ID 的值
func ContextWithAppId(ctx context.Context, appId string) context.Context {
	return Set(ctx, MetaKeyAppId, appId)
}

// AppIdFromContext 从 Context 中获取 APP ID 的值
func AppIdFromContext(ctx context.Context) string {
	return MustGet(ctx, MetaKeyAppId)
}
