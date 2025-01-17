package gcpslog_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/vvakame/sdlog/gcpslog"
	"golang.org/x/exp/slog"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ho := &gcpslog.HandlerOptions{
		Level:     slog.LevelDebug,
		ProjectID: "sdlog-test-project",
		TraceInfo: func(ctx context.Context) (string, string) {
			return "trace-id-a", "span-id-b"
		},
	}

	var buf bytes.Buffer
	h := ho.NewHandler(&buf)

	logger := slog.New(h)
	logger.Enabled(ctx, slog.LevelDebug)
	logger.InfoCtx(ctx, "info message")
	logger.ErrorCtx(ctx, "error message", "error", errors.New("error"))
	logger.LogAttrs(ctx, slog.LevelDebug, "log attrs", slog.String("key", "value"))

	t.Log(buf.String())
}

func Test_example(t *testing.T) {
	defaultLogger := slog.Default()
	t.Cleanup(func() {
		slog.SetDefault(defaultLogger)
	})

	var buf bytes.Buffer
	slog.SetDefault(slog.New(gcpslog.HandlerOptions{}.NewHandler(&buf)))

	ctx := context.Background()

	slog.InfoCtx(ctx, "info message")
	slog.ErrorCtx(ctx, "error message", "error", errors.New("error"))
	slog.LogAttrs(ctx, slog.LevelDebug, "log attrs", slog.String("key", "value"))

	t.Log(buf.String())
}
