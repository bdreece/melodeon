package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify"
	"github.com/bdreece/melodeon/pkg/view"
)

var (
	ErrOpenConfig  = errors.New("failed to open config file")
	ErrParseConfig = errors.New("failed to parse config file")
)

type Config struct {
	AppOptions

	Cookies session.Options `json:"cookies"`
	Logging logger.Options  `json:"logging"`
	Spotify spotify.Options `json:"spotify"`
	Web     view.Options    `json:"web"`
}

var defaultConfig = Config{
	AppOptions: DefaultAppOptions,
	Cookies:    session.DefaultOptions,
	Logging:    logger.DefaultOptions,
	Web:        view.DefaultOptions,
}

func Default() *Config {
	return &defaultConfig
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Join(ErrOpenConfig, err)
	}

	cfg := defaultConfig
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, errors.Join(ErrParseConfig, err)
	}

	return &cfg, nil
}
