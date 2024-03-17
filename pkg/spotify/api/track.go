package api

type Track struct {
	namedNode

	Album            SimpleAlbum  `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       struct{}     `json:"linked_from"`
	Popularity       int          `json:"popularity"`
	PreviewUrl       *string      `json:"preview_url"`
	Restrictions     Restrictions `json:"restrictions"`
	TrackNumber      int          `json:"track_number"`
	IsLocal          bool         `json:"is_local"`

	ExternalIds struct {
		Isrc string `json:"isrc"`
		Ean  string `json:"ean"`
		Upc  string `json:"upc"`
	} `json:"external_ids"`
}

