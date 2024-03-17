package spotify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type client struct {
	http.Client

	accessToken  string
	refreshToken string
	expiration   time.Time
	tokenClient  *TokenClient
}

func (client *client) Do(req *http.Request) (*http.Response, error) {
    req.Header.Add("Authorization", "Bearer "+client.accessToken)
    return client.Client.Do(req)
}

func (client *client) ensureSuccessResponse(res *http.Response) error {
    var (
        forbidden = errors.New("bad oauth request")
        tooManyRequests = errors.New("you are being rate limited")
        unauthorized = errors.New("bad or expired token")
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

func (client *client) ensureValidToken(ctx context.Context) error {
    if client.expiration.Before(time.Now()) {
        return nil
    }

    data, err := client.tokenClient.Refresh(ctx, client.refreshToken)
    if err != nil {
        return fmt.Errorf("failed to refresh token: %w", err)
    }

    client.accessToken = data.AccessToken
    client.refreshToken = data.RefreshToken
    client.expiration = time.Now().Add(time.Duration(data.ExpiresIn)*time.Second)
    return nil
}
