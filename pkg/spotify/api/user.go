package api

type User struct {
	node

	Country      string       `json:"country"`
	DisplayName  string       `json:"display_name"`
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Images       []Image      `json:"images"`
}

