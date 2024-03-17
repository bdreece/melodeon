package api

type node struct {
	Href string `json:"href"`
	Id   string `json:"id"`
	Type string `json:"type"`
	Uri  string `json:"uri"`
}

type namedNode struct {
	node

	Name string `json:"name"`
}

