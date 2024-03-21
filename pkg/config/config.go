// Package config implements configuration file deserialization
// for providing global options structures.
//
// The config package should be used during startup to inject
// options into dependencies
package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify"
	"github.com/bdreece/melodeon/pkg/store"
	"github.com/bdreece/melodeon/pkg/view"
)

var (
	ErrOpenConfig  = errors.New("failed to open config file")
	ErrParseConfig = errors.New("failed to parse config file")
)

// A Config provides the schema for configuration files
type Config struct {
	AppOptions

	Logger  logger.Options  `json:"logging"`
	Session session.Options `json:"session"`
	Spotify spotify.Options `json:"spotify"`
	Store   store.Options   `json:"store"`
	View    view.Options    `json:"web"`
}

var defaultConfig = Config{
	AppOptions: DefaultAppOptions,
	Logger:     logger.DefaultOptions,
	Session:    session.DefaultOptions,
	Store:      store.DefaultOptions,
	View:       view.DefaultOptions,
}

// Default returns the default configuration values
func Default() *Config {
	return &defaultConfig
}

// Load deserializes the configuration file located at the given path
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
