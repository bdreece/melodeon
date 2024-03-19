package spotify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bdreece/melodeon/pkg/spotify/api"
)

var ErrSearch = errors.New("failed to receive search response")

type SearchClient struct{ Client }

func (client *SearchClient) Search(ctx context.Context, req api.SearchRequest) (*api.SearchResponse, error) {
	const endpoint string = "https://accounts.spotify.com/v1/search"
	if err := client.ensureValidToken(ctx); err != nil {
		return nil, err
	}

	target, _ := url.Parse(endpoint)
	target.RawQuery = req.Query().Encode()
	r, err := http.NewRequestWithContext(ctx, "GET", target.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create search request: %w", err)
	}

	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to send search request: %w", err)
	}

	if err = client.ensureSuccessResponse(res); err != nil {
		return nil, fmt.Errorf("failed to received search response: %w", err)
	}

	data := new(api.SearchResponse)
	if err = json.NewDecoder(res.Body).Decode(data); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	return data, nil
}
