package store

type Options struct {
    Path string `json:"path"`
}

var DefaultOptions = Options{
    Path: "tmp/store.db",
}
