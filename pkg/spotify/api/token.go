package api

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

// UnmarshalJSON implements json.Unmarshaler.
func (token *Token) UnmarshalJSON(data []byte) error {
	var v struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		TokenType    string `json:"token_type"`
		Scope        string `json:"scope"`
		ExpiresIn    int    `json:"expires_in"`
	}

    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }

    token.AccessToken = v.AccessToken
    token.RefreshToken = v.RefreshToken
    token.TokenType = v.TokenType
    token.Scope = v.Scope
    token.ExpiresIn = time.Now().Add(time.Duration(v.ExpiresIn)*time.Second)
    return nil
}

var _ json.Unmarshaler = (*Token)(nil)
