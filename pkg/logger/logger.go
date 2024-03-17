package logger

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"reflect"
)

var ErrOpenFile = errors.New("failed to open log file")

func New(opts *Options) (*slog.Logger, error) {
	const perms fs.FileMode = 0o0644 // owner: rw, group: r, others: r

	f, err := os.OpenFile(opts.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perms)
	if err != nil {
		return nil, errors.Join(ErrOpenFile, err)
	}

	handler := slog.NewTextHandler(io.MultiWriter(f, os.Stdout), &slog.HandlerOptions{
		Level: slog.Level(opts.Level),
	})

	return slog.New(handler), nil
}

func For[T any](log *slog.Logger) *slog.Logger {
    t := reflect.TypeFor[T]()
    log = log.With(slog.String("context", t.String()))
    log.Debug("initializing...")

    return log
}

func Error(log *slog.Logger, msg string, err error) error {
    log.Error(msg, slog.String("error", err.Error()))
    return fmt.Errorf("%s: %w", msg, err)
}
