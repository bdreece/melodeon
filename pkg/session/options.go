package session

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
)

type Options struct {
	RootDirectory string           `json:"root_dir"`
	SecretKey     []byte           `json:"secret_key"`
	Cookies       sessions.Options `json:"cookies"`
}

var DefaultOptions = Options{
	RootDirectory: "./tmp",
	Cookies: sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	},
}

// UnmarshalJSON implements json.Unmarshaler.
func (opts *Options) UnmarshalJSON(data []byte) error {
	var v struct {
		RootDirectory *string `json:"string"`
		SecretKey     []byte  `json:"secret_key"`
		Cookies       struct {
			Path     *string `json:"path"`
			Domain   *string `json:"domain"`
			MaxAge   *int    `json:"max_age"`
			Secure   *bool   `json:"secure"`
			HttpOnly *bool   `json:"http_only"`
			SameSite *string `json:"same_site"`
		} `json:"cookies"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	opts.SecretKey = v.SecretKey
	if v.RootDirectory != nil {
		opts.RootDirectory = *v.RootDirectory
	}

	if v.Cookies.Path != nil {
		opts.Cookies.Path = *v.Cookies.Path
	}

	if v.Cookies.Domain != nil {
		opts.Cookies.Domain = *v.Cookies.Domain
	}

	if v.Cookies.MaxAge != nil {
		opts.Cookies.MaxAge = *v.Cookies.MaxAge
	}

	if v.Cookies.Secure != nil {
		opts.Cookies.Secure = *v.Cookies.Secure
	}

	if v.Cookies.HttpOnly != nil {
		opts.Cookies.HttpOnly = *v.Cookies.HttpOnly
	}

	if v.Cookies.SameSite != nil {
		switch *v.Cookies.SameSite {
		case "":
			break
		case "none":
			opts.Cookies.SameSite = http.SameSiteNoneMode
		case "strict":
			opts.Cookies.SameSite = http.SameSiteStrictMode
		case "lax":
			opts.Cookies.SameSite = http.SameSiteLaxMode
		default:
			opts.Cookies.SameSite = http.SameSiteDefaultMode
		}
	}

	return nil
}

var _ json.Unmarshaler = (*Options)(nil)
