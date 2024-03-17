package api

type SimpleAlbum struct {
	namedNode

	AlbumType            string         `json:"album_type"`
	Artists              []SimpleArtist `json:"artists"`
	AvailableMarkets     []string       `json:"available_markets"`
	ExternalUrls         ExternalUrls   `json:"external_urls"`
	Images               []Image        `json:"images"`
	ReleaseDate          string         `json:"release_date"`
	ReleaseDatePrecision string         `json:"release_date_precision"`
	Restrictions         Restrictions   `json:"restrictions"`
	TotalTracks          int            `json:"total_tracks"`
}
