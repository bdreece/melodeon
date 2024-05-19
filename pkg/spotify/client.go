package spotify

import (
	"fmt"
	"net/http"
)

const baseUri string = "https://api.spotify.com"

type Client struct {
	http.Client

	Authorizer Authorizer
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	token, err := c.Authorizer.Authorize(req.Context())
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	return c.Client.Do(req)
}
