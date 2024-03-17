package api

type Page[T any] struct {
	Href     string `json:"href"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	Total    int    `json:"total"`
	Items    []T    `json:"items"`
}


