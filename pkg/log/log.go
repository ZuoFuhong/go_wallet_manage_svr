package log

import (
	"context"
	"fmt"
	"github.com/ZuoFuhong/go_wallet_manage_svr/consts"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
)

func Debugf(format string, args ...interface{}) {
	_ = log.Output(2, fmt.Sprintf("DEBUG "+format+"\n", args...))
}

func Infof(format string, args ...interface{}) {
	_ = log.Output(2, fmt.Sprintf("INFO  "+format+"\n", args...))
}

func Warnf(format string, args ...interface{}) {
	_ = log.Output(2, fmt.Sprintf("WARN  "+format+"\n", args...))
}

func Errorf(format string, args ...interface{}) {
	_ = log.Output(2, fmt.Sprintf("ERROR "+format+"\n", args...))
}

func Fatalf(format string, args ...interface{}) {
	_ = log.Output(2, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func DebugContextf(ctx context.Context, format string, args ...interface{}) {
	traceId := metadata.ValueFromIncomingContext(ctx, consts.TraceKey)[0]
	prefix := fmt.Sprintf("DEBUG %s ", traceId)
	_ = log.Output(2, fmt.Sprintf(prefix+format+"\n", args...))
}

func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	traceId := metadata.ValueFromIncomingContext(ctx, consts.TraceKey)[0]
	prefix := fmt.Sprintf("INFO  %s ", traceId)
	_ = log.Output(2, fmt.Sprintf(prefix+format+"\n", args...))
}

func WarnContextf(ctx context.Context, format string, args ...interface{}) {
	traceId := metadata.ValueFromIncomingContext(ctx, consts.TraceKey)[0]
	prefix := fmt.Sprintf("WARN  %s ", traceId)
	_ = log.Output(2, fmt.Sprintf(prefix+format+"\n", args...))
}

func ErrorContextf(ctx context.Context, format string, args ...interface{}) {
	traceId := metadata.ValueFromIncomingContext(ctx, consts.TraceKey)[0]
	prefix := fmt.Sprintf("ERROR %s ", traceId)
	_ = log.Output(2, fmt.Sprintf(prefix+format+"\n", args...))
}
