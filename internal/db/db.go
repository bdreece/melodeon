package db

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

type Config struct {
	DataDirectory string `yaml:"data_dir"`
}

func New(cfg *Config) (*bolt.DB, error) {
	const mode fs.FileMode = 0o0644

	if err := os.MkdirAll(cfg.DataDirectory, mode); err != nil {
		return nil, err
	}

	path := filepath.Join(cfg.DataDirectory, "melodeon.db")

	return bolt.Open(path, mode, nil)
}
