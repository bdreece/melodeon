package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const searchMarket string = "US"

var searchTypes string

func init() {
	types, _ := json.Marshal([]string{
		"tracks",
		"artists",
		"albums",
		"playlists",
	})

	searchTypes = string(types)
}

func (req SearchRequest) Query() url.Values {
	return url.Values{
		"q":      {req.Q},
		"market": {searchMarket},
		"types":  {searchTypes},
		"limit":  {fmt.Sprint(req.Limit)},
		"offset": {fmt.Sprint(req.Offset)},
	}
}
