package logging

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"strings"
	"sync"
)

type JSONDebugHandler struct {
	opts DebugHandlerOptions
	mu   *sync.Mutex
	out  io.Writer
}

type DebugHandlerOptions struct {
	// Level はログに出力する最小限のレベルを指定する
	// このレベルよりも下のレベルのログは出力されない
	// デフォルトでは slog.LevelInfo を使用する
	Level slog.Leveler

	// Indent は出力されるJSONログのインデントを指定する
	// 指定しない場合、1より小さい値を指定した場合は 2 とする
	Indent int
}

func NewJSONDebugHandler(out io.Writer, opts *DebugHandlerOptions) *JSONDebugHandler {
	h := &JSONDebugHandler{out: out, mu: &sync.Mutex{}}
	if opts.Level != nil {
		h.opts.Level = opts.Level
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	if opts.Indent != 0 {
		h.opts.Indent = opts.Indent
	}
	if h.opts.Indent <= 0 {
		h.opts.Indent = 2
	}
	return h
}

// Enabled は引数を処理する前に呼ばれ、処理を続行するかを確認する
func (h *JSONDebugHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *JSONDebugHandler) Handle(_ context.Context, r slog.Record) error {
	fields := make(map[string]any, r.NumAttrs()+3)
	fields["time"] = r.Time
	fields["level"] = r.Level
	fields["message"] = r.Message

	r.Attrs(func(a slog.Attr) bool {
		addFields(fields, a)
		return true
	})

	data, err := json.MarshalIndent(fields, "", strings.Repeat(" ", h.opts.Indent))
	if err != nil {
		return err
	}
	data = append(data, '\n')

	h.out.Write(data)
	return nil
}

func addFields(fields map[string]any, a slog.Attr) {
	value := a.Value.Any()
	attrs, ok := value.([]slog.Attr)
	if !ok {
		fields[a.Key] = value
		return
	}

	// ネストしている場合、再起的にフィールドを探索する
	innerFields := make(map[string]any, len(attrs))
	for _, attr := range attrs {
		addFields(innerFields, attr)
	}
	fields[a.Key] = innerFields
}

func (h *JSONDebugHandler) WithGroup(name string) slog.Handler {
	return h
}

func (h *JSONDebugHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}
