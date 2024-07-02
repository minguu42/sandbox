package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

type JSONIndentHandler struct {
	w       io.Writer
	opts    *slog.HandlerOptions
	handler slog.Handler
	buf     *bytes.Buffer
}

func NewJSONIndentHandler(w io.Writer, opts *slog.HandlerOptions) *JSONIndentHandler {
	buf := &bytes.Buffer{}
	return &JSONIndentHandler{
		w:       w,
		opts:    opts,
		handler: slog.NewJSONHandler(buf, opts),
		buf:     buf,
	}
}

func (h *JSONIndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *JSONIndentHandler) Handle(ctx context.Context, record slog.Record) error {
	if err := h.handler.Handle(ctx, record); err != nil {
		return err
	}

	encoder := json.NewEncoder(h.w)
	encoder.SetIndent("", strings.Repeat(" ", 2))
	if err := encoder.Encode(json.RawMessage(h.buf.Bytes())); err != nil {
		return fmt.Errorf("failed to encode json log entry: %w", err)
	}

	h.buf.Reset()
	return nil
}

func (h *JSONIndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	buf := &bytes.Buffer{}
	return &JSONIndentHandler{
		w:       h.w,
		opts:    h.opts,
		handler: slog.NewJSONHandler(buf, h.opts).WithAttrs(attrs),
		buf:     buf,
	}
}

func (h *JSONIndentHandler) WithGroup(name string) slog.Handler {
	buf := &bytes.Buffer{}
	return &JSONIndentHandler{
		w:       h.w,
		opts:    h.opts,
		handler: slog.NewJSONHandler(buf, h.opts).WithGroup(name),
		buf:     buf,
	}
}
