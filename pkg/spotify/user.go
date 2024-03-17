package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserClient struct { client }

func (client *UserClient) GetCurrentUser(ctx context.Context) (*User, error) {
    const endpoint string = "https://accounts.spotify.com/v1/me"
    if err := client.ensureValidToken(ctx); err != nil {
        return nil, err
    }

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create curent user request: %w", err)
    }

    res, err := client.Do(req)    
    if err != nil {
        return nil, fmt.Errorf("failed to send current user request: %w", err)
    }

    if err = client.ensureSuccessResponse(res); err != nil {
        return nil, fmt.Errorf("failed to received search response: %w", err)
    }

    data := new(User)
    if err = json.NewDecoder(res.Body).Decode(data); err != nil {
        return nil, fmt.Errorf("failed to decode user response: %w", err)
    }

    return data, nil
}

func NewUserClient(opts *ClientOptions) *UserClient {
    return &UserClient{
        client: client{
            accessToken: opts.AccessToken,
            refreshToken: opts.RefreshToken,
            expiration: opts.Expiration,
        },
    }
}
