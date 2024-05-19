package db

import (
	"io/fs"

	"github.com/boltdb/bolt"
)

type Config struct {
	Path string `yaml:"path"`
}

func New(cfg *Config) (*bolt.DB, error) {
	const mode fs.FileMode = 0o0644

	return bolt.Open(cfg.Path, mode, nil)
}
