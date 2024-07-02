package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
)

type JSONDebugHandler2 struct {
	w    io.Writer
	opts *slog.HandlerOptions

	handler slog.Handler
	buf     *bytes.Buffer
}

func NewJSONDebugHandler2(w io.Writer, opts *slog.HandlerOptions) *JSONDebugHandler2 {
	buf := &bytes.Buffer{}
	return &JSONDebugHandler2{
		w:       w,
		opts:    opts,
		handler: slog.NewJSONHandler(buf, opts),
		buf:     buf,
	}
}

func (h *JSONDebugHandler2) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *JSONDebugHandler2) WithAttrs(attrs []slog.Attr) slog.Handler {
	buf := &bytes.Buffer{}
	return &JSONDebugHandler2{
		w:       h.w,
		opts:    h.opts,
		handler: slog.NewJSONHandler(buf, h.opts).WithAttrs(attrs),
		buf:     buf,
	}
}

func (h *JSONDebugHandler2) WithGroup(name string) slog.Handler {
	buf := &bytes.Buffer{}
	return &JSONDebugHandler2{
		w:       h.w,
		opts:    h.opts,
		handler: slog.NewJSONHandler(buf, h.opts).WithGroup(name),
		buf:     buf,
	}
}

func (h *JSONDebugHandler2) Handle(ctx context.Context, r slog.Record) error {
	if err := h.handler.Handle(ctx, r); err != nil {
		return err
	}

	encoder := json.NewEncoder(h.w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(json.RawMessage(h.buf.Bytes())); err != nil {
		return fmt.Errorf("failed to encode json log entry: %w", err)
	}

	h.buf.Reset()
	return nil
}
