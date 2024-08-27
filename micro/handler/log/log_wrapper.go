package log

import (
	"context"
	"time"

	"github.com/jinmukeji/plat-pkg/v4/micro/errors"
	"github.com/jinmukeji/plat-pkg/v4/micro/meta"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
)

const (
	logCidKey     = "cid"
	logLatencyKey = "latency"
	logRpcCallKey = "rpc.call"

	// rpcMetadata = "[RPC METADATA]"
	rpcFailed = "[RPC ERR]"
	rpcOk     = "[RPC OK]"

	errorField = "error"
	// errorCodeField = "errcode"
)

var defaultHelperLogger *logger.Helper

func helperLogger() *logger.Helper {
	if defaultHelperLogger == nil {
		return logger.NewHelper(logger.DefaultLogger)
	}
	return defaultHelperLogger
}

// LogWrapper is a handler wrapper that server request
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		start := time.Now()
		cid := meta.CidFromContext(ctx)

		// 注入一个包含 cid Field 的 logger.Entry
		hl := helperLogger()
		cl := hl.WithFields(map[string]interface{}{logCidKey: cid})
		c := logger.NewContext(ctx, logger.DefaultLogger)

		err := fn(c, req, rsp)
		// RPC 计算经历的时间长度
		end := time.Now()
		latency := end.Sub(start)

		l := cl.WithFields(map[string]interface{}{
			logRpcCallKey: req.Method(),
			logLatencyKey: latency.String(),
		})

		switch v := err.(type) {
		case nil:
			l.Info(rpcOk)
		case *errors.RpcError:
			l.WithFields(map[string]interface{}{
				errorField: v.DetailedError(),
			}).Warn(rpcFailed)
		case error:
			l.WithFields(map[string]interface{}{errorField: err.Error()}).Warn(rpcFailed)
		default:
			l.Errorf("unknown error type: %v", v)
		}

		return err
	}
}
