package trace

import (
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gofrs/uuid"
)

type Config struct {
	Directory string `yaml:"dir"`
	Level     int    `yaml:"level"`
}

type Logger struct {
	slog.Logger
}

func (log *Logger) Trace() *Trace {
	id, _ := uuid.NewV4()

	return &Trace{
		Logger: *log.With(slog.String("trace", id.String())),
		ID:     id,
	}
}

func NewLogger(cfg *Config) (*Logger, error) {
	const (
		flag int         = os.O_CREATE | os.O_APPEND | os.O_WRONLY
		mode fs.FileMode = 0o0644
	)

	path := filepath.Join(cfg.Directory, "melodeon.log")
	f, err := os.OpenFile(path, flag, mode)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file %q: %v", path, err)
	}

	w := io.MultiWriter(f, os.Stdout)
	log := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.Level(cfg.Level),
	}))

	return &Logger{*log}, nil
}
