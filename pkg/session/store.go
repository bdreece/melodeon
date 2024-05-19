package session

import (
	"github.com/gorilla/sessions"
)

// NewStore creates a new cookie session store.
func NewStore(cfg *Config) *sessions.CookieStore {
	store := sessions.NewCookieStore(cfg.SigningKey, cfg.EncryptionKey)
	store.Options = cfg.Cookie.SessionOptions()
	return store
}
