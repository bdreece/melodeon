package spotify

import (
	"encoding/json"
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresIn    time.Time
	Scope        string
}

func (t *Token) UnmarshalJSON(b []byte) error {
	value := new(struct {
		AccessToken  string `json:"access_token" mapstructure:"access_token"`
		RefreshToken string `json:"refresh_token" mapstructure:"refresh_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in" mapstructure:"expiration"`
		Scope        string `json:"scope"`
	})

	if err := json.Unmarshal(b, value); err != nil {
		return err
	}

	t.AccessToken = value.AccessToken
	t.RefreshToken = value.RefreshToken
	t.TokenType = value.TokenType
	t.Scope = value.Scope
	t.ExpiresIn = time.Now().Add(time.Duration(value.ExpiresIn) * time.Second)

	return nil
}
