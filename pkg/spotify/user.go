package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

type (
	TopItemsType interface{ Track | Artist }

	TopItemsRequest struct {
		TimeRange string
		Limit     int
		Offset    int
	}

	TopItemsResponse[T TopItemsType] struct {
		Href     string `json:"href"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Limit    int    `json:"limit"`
		Offset   int    `json:"offset"`
		Total    int    `json:"total"`
		Items    []T    `json:"items"`
	}
)

type UserClient struct{ Client }

var (
	topArtistsUri, _   = url.Parse(baseUri + "/v1/me/top/artists")
	topTracksUri, _    = url.Parse(baseUri + "/v1/me/top/tracks")
	topItemsTrackType  = reflect.TypeFor[Track]().Name()
	topItemsArtistType = reflect.TypeFor[Artist]().Name()
)

func (c *UserClient) Profile(ctx context.Context) (*User, error) {
	const uri string = baseUri + "/v1/me"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	user := new(User)
	if err = json.NewDecoder(res.Body).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UserClient) TopTracks(ctx context.Context, req TopItemsRequest) (*TopItemsResponse[Track], error) {
	return userTopItems[Track](ctx, c, *topTracksUri, req)
}

func (c *UserClient) TopArtists(ctx context.Context, req TopItemsRequest) (*TopItemsResponse[Artist], error) {
	return userTopItems[Artist](ctx, c, *topArtistsUri, req)
}

func userTopItems[T TopItemsType](
	ctx context.Context,
	client *UserClient,
	uri url.URL,
	req TopItemsRequest,
) (*TopItemsResponse[T], error) {
	var itemType string

	t := reflect.TypeFor[T]()
	switch t.Name() {
	case topItemsTrackType:
		itemType = "tracks"
	case topItemsArtistType:
		itemType = "artists"
	default:
		panic("invalid TopItemsType")
	}

	qs := url.Values{
		"type":       {itemType},
		"time_range": {req.TimeRange},
		"limit":      {fmt.Sprint(req.Limit)},
		"offset":     {fmt.Sprint(req.Offset)},
	}

	uri.RawQuery = qs.Encode()
	httpreq, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	httpres, err := client.Do(httpreq)
	if err != nil {
		return nil, err
	}

	defer httpres.Body.Close()
	res := new(TopItemsResponse[T])
	if err = json.NewDecoder(httpres.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}
