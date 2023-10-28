package logs

import "context"

func CtxFatal(ctx context.Context, format string, v ...interface{}) {
}

func CtxError(ctx context.Context, format string, v ...interface{}) {
}

func CtxWarn(ctx context.Context, format string, v ...interface{}) {
}

func CtxNotice(ctx context.Context, format string, v ...interface{}) {
}

func CtxInfo(ctx context.Context, format string, v ...interface{}) {
}
