package logging

import (
	"context"
	"log/slog"
	"os"
)

var applicationLogger *slog.Logger

func init() {
	applicationLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "message"
			}
			return a
		},
	}))
}

func InitDebug() {
	applicationLogger = slog.New(NewJSONDebugHandler(os.Stdout, &DebugHandlerOptions{
		Indent: 2,
	}))
}

type loggerKey struct{}

// logger はコンテキストからリクエストロガーを取り出す
// コンテキストにリクエストロガーが存在しなければアプリケーションロガーを使用する
func logger(ctx context.Context) *slog.Logger {
	v, ok := ctx.Value(loggerKey{}).(*slog.Logger)
	if ok {
		return v
	}
	return applicationLogger
}

// ContextWithLogger はリクエストロガーを生成し、コンテキストにリクエストロガーをセットする
func ContextWithLogger(ctx context.Context, method string) context.Context {
	l := applicationLogger.With(slog.String("request_id", "1234567890"))
	return context.WithValue(ctx, loggerKey{}, l)
}

func LogExample(ctx context.Context) {
	logger(ctx).Debug("debug")
	logger(ctx).Info("info")
	logger(ctx).Warn("warn")
	logger(ctx).Error("error")

	logger(ctx).LogAttrs(ctx, slog.LevelInfo, "foo bar baz",
		slog.Duration("execution_time", 456),
		slog.Group("request",
			slog.String("id", "some-long-id"),
			slog.Int("content_length", 100000),
		),
	)
}
