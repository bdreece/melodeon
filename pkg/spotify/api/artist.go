package api

type Artist struct {
	namedNode

	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Genres       []string     `json:"genres"`
	Images       []Image      `json:"images"`
	Popularity   int          `json:"popularity"`
}

type SimpleArtist struct {
	namedNode

	ExternalUrls ExternalUrls `json:"external_urls"`
}

