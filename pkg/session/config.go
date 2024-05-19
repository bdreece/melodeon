package session

import (
	"encoding/base64"
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/yaml.v3"
)

type Key []byte

func (k *Key) UnmarshalYAML(node *yaml.Node) error {
	value, err := base64.StdEncoding.DecodeString(node.Value)
	if err != nil {
		return err
	}

	*k = value
	return nil
}

type Config struct {
	SigningKey    Key          `yaml:"signing_key"`
	EncryptionKey Key          `yaml:"encrypt_key"`
	Cookie        CookieConfig `yaml:"cookie"`
}

type CookieConfig struct {
	Path     string `yaml:"path"`
	Domain   string `yaml:"domain"`
	MaxAge   int    `yaml:"max_age"`
	Secure   bool   `yaml:"secure"`
	HttpOnly bool   `yaml:"http_only"`
	SameSite string `yaml:"same_site"`
}

func (c CookieConfig) SessionOptions() *sessions.Options {
	var sameSite http.SameSite

	switch c.SameSite {
	case "none":
		sameSite = http.SameSiteNoneMode
	case "lax":
		sameSite = http.SameSiteLaxMode
	case "strict":
		sameSite = http.SameSiteStrictMode
	default:
		sameSite = http.SameSiteDefaultMode
	}

	return &sessions.Options{
		Path:     c.Path,
		Domain:   c.Domain,
		MaxAge:   c.MaxAge,
		Secure:   c.Secure,
		HttpOnly: c.HttpOnly,
		SameSite: sameSite,
	}
}
