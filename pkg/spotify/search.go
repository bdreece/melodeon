package spotify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var ErrSearch = errors.New("failed to receive search response")

type SearchClient struct{ client }

func (client *SearchClient) Search(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
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

	data := new(SearchResponse)
	if err = json.NewDecoder(res.Body).Decode(data); err != nil {
        return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

    return data, nil
}

func NewSearchClient(opts *ClientOptions) *SearchClient {
    return &SearchClient{
        client: client{
            accessToken: opts.AccessToken,
            refreshToken: opts.RefreshToken,
            expiration: opts.Expiration,
        },
    }
}

type SearchRequest struct {
	Q      string
	Limit  int
	Offset int
}

func (req SearchRequest) Query() url.Values {
	const market string = "US"
	types, _ := json.Marshal([]string{
		"tracks",
		"artists",
		"albums",
		"playlists",
	})

	return url.Values{
		"q":      {req.Q},
		"market": {market},
		"types":  {string(types)},
		"limit":  {fmt.Sprint(req.Limit)},
		"offset": {fmt.Sprint(req.Offset)},
	}
}

type SearchResponse struct {
	Tracks    Page[Track]          `json:"tracks"`
	Artists   Page[Artist]         `json:"artists"`
	Albums    Page[SimpleAlbum]    `json:"albums"`
	Playlists Page[SimplePlaylist] `json:"playlists"`
}
