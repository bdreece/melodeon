package spotify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/bdreece/melodeon/pkg/spotify/api"
)

type Client struct {
	http.Client
	api.Token

	handler *TokenHandler
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+client.AccessToken)
	return client.Client.Do(req)
}

func (client *Client) ensureSuccessResponse(res *http.Response) error {
	var (
		forbidden       = errors.New("bad oauth request")
		tooManyRequests = errors.New("you are being rate limited")
		unauthorized    = errors.New("bad or expired token")
	)

	switch res.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return unauthorized
	case http.StatusForbidden:
		return forbidden
	case http.StatusTooManyRequests:
		return tooManyRequests
	default:
		return fmt.Errorf("unknown status code: %s", res.Status)
	}
}

func (client *Client) ensureValidToken(ctx context.Context) error {
	if client.ExpiresIn.Before(time.Now()) {
		return nil
	}

	data, err := client.handler.Refresh(ctx, client.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	client.Token = *data
	return nil
}

func NewClient(token *api.Token, handler *TokenHandler) *Client {
	return &Client{
		Token:   *token,
		handler: handler,
	}
}
