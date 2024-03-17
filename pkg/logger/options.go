package logger

type Options struct {
	File  string `json:"file"`
	Level int    `json:"level"`
}

var DefaultOptions = Options{
	File:  "./tmp/melodeon.log",
	Level: -4,
}
