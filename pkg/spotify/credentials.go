package spotify

import (
	"encoding/base64"
	"net/url"
)

type Credentials struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

func (c Credentials) Userinfo() string {
	userinfo := url.UserPassword(c.ClientID, c.ClientSecret).String()
	return base64.StdEncoding.EncodeToString([]byte(userinfo))
}
