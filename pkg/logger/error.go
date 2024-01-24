package logger

import (
	"context"
	"fmt"
	"log/slog"
)

type SlogError struct {
	err   error
	msg   string
	attrs []slog.Attr
}

func (e *SlogError) Error() string {
	if e.err == nil {
		return e.msg
	}
	return fmt.Sprintf("%s: %s", e.msg, e.err.Error())
}

func (e *SlogError) Log(ctx context.Context) {
	if e.err != nil {
		e.attrs = append(e.attrs, slog.String("err", e.err.Error()))
	}
	slog.LogAttrs(ctx, slog.LevelError, e.msg, e.attrs...)
}

func NewSlogError(err error, msg string, attrs ...slog.Attr) *SlogError {
	return &SlogError{
		err:   err,
		msg:   msg,
		attrs: attrs,
	}
}
