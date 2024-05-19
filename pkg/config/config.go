package config

import (
	"os"

	"github.com/bdreece/melodeon/internal/db"
	"github.com/bdreece/melodeon/internal/renderer"
	"github.com/bdreece/melodeon/internal/trace"
	"github.com/bdreece/melodeon/pkg/rooms"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify"
	"gopkg.in/yaml.v3"
)

type Config struct {
	RootConfig `yaml:",inline"`

	DB       db.Config      `yaml:"db"`
	Logging  trace.Config   `yaml:"logging"`
	Rooms    rooms.Config   `yaml:"rooms"`
	Sessions session.Config `yaml:"sessions"`
	Spotify  spotify.Config `yaml:"spotify"`
	Web      WebConfig      `yaml:"web"`
}

type WebConfig struct {
	renderer.Config `yaml:",inline"`

	StaticDirectory string `yaml:"static_dir"`
	AssetsDirectory string `yaml:"assets_dir"`
}

type RootConfig struct {
	Mode Mode `yaml:"mode"`
}

func Parse(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	cfg := new(Config)
	if err = yaml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
