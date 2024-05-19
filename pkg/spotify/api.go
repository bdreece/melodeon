package spotify

type (
	Node struct {
		Href string `json:"href"`
		Id   string `json:"id"`
		Type string `json:"type"`
		Uri  string `json:"uri"`
	}

	NamedNode struct {
		Node

		Name string `json:"name"`
	}

	SimpleAlbum struct {
		NamedNode

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

	Artist struct {
		NamedNode

		ExternalUrls ExternalUrls `json:"external_urls"`
		Followers    Followers    `json:"followers"`
		Genres       []string     `json:"genres"`
		Images       []Image      `json:"images"`
		Popularity   int          `json:"popularity"`
	}

	SimpleArtist struct {
		NamedNode

		ExternalUrls ExternalUrls `json:"external_urls"`
	}

	ExternalUrls struct {
		Spotify string `json:"spotify"`
	}

	Followers struct {
		Href  *string `json:"href"`
		Total int     `json:"total"`
	}

	Image struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	}

	Page[T any] struct {
		Href     string `json:"href"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Limit    int    `json:"limit"`
		Offset   int    `json:"offset"`
		Total    int    `json:"total"`
		Items    []T    `json:"items"`
	}

	SimplePlaylist struct {
		NamedNode

		Collaborative bool         `json:"collaborative"`
		Description   string       `json:"description"`
		ExternalUrls  ExternalUrls `json:"external_urls"`
		Images        []Image      `json:"images"`
		Public        bool         `json:"public"`
		SnapshotId    string       `json:"snapshot_id"`

		Owner struct {
			Node
			ExternalUrls ExternalUrls `json:"external_urls"`
			Followers    Followers    `json:"followers"`
			DisplayName  *string      `json:"display_name"`
		} `json:"owner"`

		Tracks struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"tracks"`
	}

	Restrictions struct {
		Reason string `json:"reason"`
	}

	Track struct {
		NamedNode

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

	User struct {
		Node

		Country      string       `json:"country"`
		DisplayName  string       `json:"display_name"`
		ExternalUrls ExternalUrls `json:"external_urls"`
		Followers    Followers    `json:"followers"`
		Images       []Image      `json:"images"`
	}
)
