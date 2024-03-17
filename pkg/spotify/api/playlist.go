package api

type SimplePlaylist struct {
	namedNode

	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalUrls  ExternalUrls `json:"external_urls"`
	Images        []Image      `json:"images"`
	Public        bool         `json:"public"`
	SnapshotId    string       `json:"snapshot_id"`

	Owner struct {
		node
		ExternalUrls ExternalUrls `json:"external_urls"`
		Followers    Followers    `json:"followers"`
		DisplayName  *string      `json:"display_name"`
	} `json:"owner"`

	Tracks struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"tracks"`
}

