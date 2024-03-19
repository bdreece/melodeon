package spotify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/bdreece/melodeon/pkg/spotify/api"
)

var ErrTokenRequestFailed = errors.New("token request failed")

type TokenHandler struct {
	http.Client

	userinfo    string
	redirectURI string
}

func (handler *TokenHandler) AuthorizeServer(ctx context.Context) (*api.Token, error) {
	return handler.send(ctx, &url.Values{
		"grant_type": {"client_credentials"},
	})
}

func (handler *TokenHandler) ExchangeCode(ctx context.Context, code string) (*api.Token, error) {
	return handler.send(ctx, &url.Values{
		"grant_type":   {"authorization_code"},
		"redirect_uri": {handler.redirectURI},
		"code":         {code},
	})
}

func (handler *TokenHandler) Refresh(ctx context.Context, refreshToken string) (*api.Token, error) {
	return handler.send(ctx, &url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	})
}

func (handler *TokenHandler) send(ctx context.Context, form *url.Values) (*api.Token, error) {
	const endpoint string = "https://accounts.spotify.com/api/token"
	body := strings.NewReader(form.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Authorization", "Basic "+handler.userinfo)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := handler.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send token request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received invalid status code %d: %w",
			res.StatusCode, ErrTokenRequestFailed)
	}

	data := new(api.Token)
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	return data, nil
}

func NewTokenHandler(opts *Options) *TokenHandler {
	return &TokenHandler{
		redirectURI: opts.RedirectURI,
		userinfo: base64.StdEncoding.EncodeToString(
			[]byte(opts.ClientID + ":" + opts.ClientSecret)),
	}
}
