package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type (
	SearchRequest struct {
		Query  string
		Limit  int
		Offset int
	}

	SearchResponse struct {
		Tracks    Page[Track]          `json:"tracks"`
		Artists   Page[Artist]         `json:"artists"`
		Albums    Page[SimpleAlbum]    `json:"albums"`
		Playlists Page[SimplePlaylist] `json:"playlists"`
	}

	SearchClient struct{ Client }
)

var searchUri, _ = url.Parse(baseUri + "/v1/search")

func (c *SearchClient) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	const (
		market string = "US"
		types  string = "album,artist,playlist,track"
	)

	qs := url.Values{
		"q":      {req.Query},
		"type":   {types},
		"market": {market},
		"limit":  {fmt.Sprint(req.Limit)},
		"offset": {fmt.Sprint(req.Offset)},
	}

	uri := *searchUri
	uri.RawQuery = qs.Encode()
	httpreq, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	httpres, err := c.Do(httpreq)
	if err != nil {
		return nil, err
	}

	defer httpres.Body.Close()
	res := new(SearchResponse)
	if err = json.NewDecoder(httpres.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}
