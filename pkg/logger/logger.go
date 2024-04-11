package logger

import (
	"context"
	"log/slog"
	"os"
)

type LevelHandler struct {
	level   slog.Leveler
	handler slog.Handler
}

func NewLevelHandler(level slog.Leveler, h slog.Handler) *LevelHandler {
	if lh, ok := h.(*LevelHandler); ok {
		h = lh.Handler()
	}
	return &LevelHandler{level, h}
}

func (h *LevelHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

func (h *LevelHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.handler.Handle(ctx, r)
}

func (h *LevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithAttrs(attrs))
}

func (h *LevelHandler) WithGroup(name string) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithGroup(name))
}

func (h *LevelHandler) Handler() slog.Handler {
	return h.handler
}

var Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

// func main() {
// th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime})
// logger := slog.New(NewLevelHandler(slog.LevelWarn, th))
// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
// logger.Info("not printed")
// logger.Warn("printed")
// loggir := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime}))
// loggir.Info("hello, world", "user", os.Getenv("USER"))

// }
