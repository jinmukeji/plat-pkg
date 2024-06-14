package microerr

import (
	"context"
	"strings"

	"github.com/jinmukeji/plat-pkg/v4/micro/errors"
	"go-micro.dev/v4"
	me "go-micro.dev/v4/errors"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
)

// 错误样式
const (
	// ErrStyleSimple 简易模式错误输出样式。通常用于面向最终用户输出，本样式隐藏内部错误信息。
	ErrStyleSimple = "microsimple"
	// ErrStyleDetailed 详细模式错误输出样式。通常用于面向开发者或系统运维人员输出，本样式输出内部错误信息。
	ErrStyleDetailed = "microdetailed"
	// ErrStyleRaw Micro原始错误输出样式。
	ErrStyleRaw = "raw"
)

const (
	ErrorStyleMetaKey      = "x-err-style"       // Metadata key / HTTP Header
	ApplicationErrorStatus = "Application Error" // 错误状态
	ApplicationErrorCode   = 520                 // 错误状态码
)

func MicroErrWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		err := fn(ctx, req, rsp)
		if err != nil {
			var rErr error

			style := errorStyleFromContext(ctx)
			rErr = err

			switch style {
			case ErrStyleSimple:
				// 使用简化版信息输出，不输出内部错误信息
				if e, ok := err.(*errors.RpcError); ok {
					rErr = wrapError(ctx, e.Error())
				}

			case ErrStyleDetailed:
				// 使用详细版信息输出，输出内部错误信息
				if e, ok := err.(*errors.RpcError); ok {
					rErr = wrapError(ctx, e.DetailedError())
				}
			case ErrStyleRaw:
				// 不做处理，直接输出原始的
			default:
				// 其它未知的，直接输出原始的
			}

			return rErr
		}

		return nil
	}
}

func wrapError(ctx context.Context, detail string) error {
	var svcName string
	if svc, ok := micro.FromContext(ctx); ok {
		svcName = svc.Server().Options().Name
	}

	err := &me.Error{
		Id:     svcName,
		Code:   ApplicationErrorCode,
		Detail: detail,
		Status: ApplicationErrorStatus,
	}

	return err
}

func errorStyleFromContext(ctx context.Context) string {
	style := ErrStyleRaw

	if md, ok := metadata.FromContext(ctx); ok {
		style = md[ErrorStyleMetaKey]
		style = strings.ToLower(style)
	}

	return style
}
