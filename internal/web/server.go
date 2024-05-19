package web

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/bdreece/melodeon/pkg/config"
)

type Server interface {
	Serve(net.Listener, http.Handler) error
}

type ServerFunc func(net.Listener, http.Handler) error

func (fn ServerFunc) Serve(l net.Listener, h http.Handler) error { return fn(l, h) }

func NewServer(cfg *config.RootConfig) (Server, error) {
	switch cfg.Mode {
	case config.HTTP:
		return ServerFunc(http.Serve), nil
	case config.FCGI:
		return ServerFunc(fcgi.Serve), nil
	default:
		return nil, fmt.Errorf("invalid mode: %q", cfg.Mode)
	}
}
