package session

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
)

type Options struct {
	sessions.Options

	SecretKey []byte
}

var DefaultOptions = Options{
    Options: sessions.Options{
        Path: "/",
        MaxAge: 0,
        HttpOnly: true,
        Secure: false,
        SameSite: http.SameSiteStrictMode,
    },
}

// UnmarshalJSON implements json.Unmarshaler.
func (opts *Options) UnmarshalJSON(data []byte) error {
	var v struct {
		Path      string `json:"path"`
		Domain    string `json:"domain"`
		MaxAge    *int    `json:"max_age"`
		Secure    *bool   `json:"secure"`
		HttpOnly  *bool   `json:"http_only"`
		SameSite  string `json:"same_site"`
		SecretKey []byte `json:"secret_key"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

    if v.Path != "" {
	    opts.Path = v.Path
    }

    if v.Domain != "" {
        opts.Domain = v.Domain
    }

    if v.MaxAge != nil {
	    opts.MaxAge = *v.MaxAge
    }

    if v.Secure != nil {
	    opts.Secure = *v.Secure
    }

    if v.HttpOnly != nil {
	    opts.HttpOnly = *v.HttpOnly
    }

	opts.SecretKey = v.SecretKey

	switch v.SameSite {
    case "":
        break
	case "none":
		opts.SameSite = http.SameSiteNoneMode
	case "strict":
		opts.SameSite = http.SameSiteStrictMode
	case "lax":
		opts.SameSite = http.SameSiteLaxMode
	default:
		opts.SameSite = http.SameSiteDefaultMode
	}

	return nil
}

var _ json.Unmarshaler = (*Options)(nil)
