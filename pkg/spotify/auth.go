package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AuthManager struct {
	config *Config
}

type Authorizer interface {
	Token() Token
	SetToken(token Token)
	Authorize(context.Context) (*Token, error)
}

type baseAuthorizer struct {
	http.Client

	token *Token
	creds Credentials
}

type userAuthorizer struct {
	baseAuthorizer

	code        string
	redirectURI string
}

type serverAuthorizer struct {
	baseAuthorizer
}

func (manager *AuthManager) User(code string) Authorizer {
	return &userAuthorizer{
		code:        code,
		redirectURI: manager.config.RedirectURI,
		baseAuthorizer: baseAuthorizer{
			creds: manager.config.Credentials,
		},
	}
}

func (manager *AuthManager) Server() Authorizer {
	return &serverAuthorizer{
		baseAuthorizer: baseAuthorizer{
			creds: manager.config.Credentials,
		},
	}
}

func (a userAuthorizer) RedirectURI() string { return a.redirectURI }

func (a *userAuthorizer) Authorize(ctx context.Context) (*Token, error) {
	if a.token == nil {
		return a.authorize(ctx)
	}

	if a.token.ExpiresIn.Compare(time.Now()) <= 0 {
		return a.refresh(ctx)
	}

	return a.token, nil

}

func (a *userAuthorizer) authorize(ctx context.Context) (*Token, error) {
	var (
		form = url.Values{
			"grant_type":   {"authorization_code"},
			"code":         {a.code},
			"redirect_uri": {a.redirectURI},
		}

		headers = http.Header{
			"Authorization": {fmt.Sprintf("Basic %s", a.creds.Userinfo())},
			"Content-Type":  {"application/x-www-form-urlencoded"},
		}
	)

	return a.requestToken(ctx, form, headers)
}

func (a *userAuthorizer) refresh(ctx context.Context) (*Token, error) {
	var (
		form = url.Values{
			"grant_type":    {"refresh_token"},
			"refresh_token": {a.token.RefreshToken},
		}

		headers = http.Header{
			"Authorization": {fmt.Sprintf("Basic %s", a.creds.Userinfo())},
			"Content-Type":  {"application/x-www-form-urlencoded"},
		}
	)

	return a.requestToken(ctx, form, headers)
}

func (a *serverAuthorizer) Authorize(ctx context.Context) (*Token, error) {
	if a.token == nil {
		return a.authorize(ctx)
	}

	if a.token.ExpiresIn.Compare(time.Now()) <= 0 {
		return a.refresh(ctx)
	}

	return a.token, nil
}

func (a *serverAuthorizer) authorize(ctx context.Context) (*Token, error) {
	var (
		form = url.Values{
			"grant_type": {"client_credentials"},
		}

		headers = http.Header{
			"Authorization": {fmt.Sprintf("Basic %s", a.creds.Userinfo())},
			"Content-Type":  {"application/x-www-form-urlencoded"},
		}
	)

	return a.requestToken(ctx, form, headers)
}

func (a *serverAuthorizer) refresh(ctx context.Context) (*Token, error) {
	var (
		form = url.Values{
			"grant_type":    {"refresh_token"},
			"refresh_token": {a.token.RefreshToken},
		}

		headers = http.Header{
			"Content-Type": {"application/x-www-form-urlencoded"},
		}
	)

	return a.requestToken(ctx, form, headers)
}

func (a baseAuthorizer) Token() Token             { return *a.token }
func (a baseAuthorizer) Credentials() Credentials { return a.creds }

func (a *baseAuthorizer) SetToken(token Token) {
	a.token = &token
}

func (a *baseAuthorizer) requestToken(
	ctx context.Context,
	form url.Values,
	headers http.Header,
) (*Token, error) {
	const uri string = "https://accounts.spotify.com/api/token"
	body := strings.NewReader(form.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}

	maps.Copy(req.Header, headers)
	res, err := a.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received invalid status code: %s", res.Status)
	}

	a.token = new(Token)
	if err = json.NewDecoder(res.Body).Decode(a.token); err != nil {
		return nil, err
	}

	return a.token, nil
}

func NewAuthManager(cfg *Config) *AuthManager {
	return &AuthManager{cfg}
}
