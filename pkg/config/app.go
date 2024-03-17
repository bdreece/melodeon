package config

import "fmt"

type AppOptions struct {
	Mode Mode `json:"mode"`
	Port int  `json:"port"`
}

func (opts AppOptions) Addr() string {
    return fmt.Sprintf(":%d", opts.Port)
}

var DefaultAppOptions = AppOptions{
    Mode: HTTP,
    Port: 3000,
}

