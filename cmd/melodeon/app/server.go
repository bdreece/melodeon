package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/bdreece/melodeon/pkg/config"
)

type Server struct {
	handler  http.Handler
	listener net.Listener
	serve    func(net.Listener, http.Handler) error
}

func (srv Server) Addr() string { return srv.listener.Addr().String() }

func (srv *Server) Serve(ctx context.Context) error {
	errch := make(chan error, 1)
	defer close(errch)

	go func() {
		errch <- srv.serve(srv.listener, srv.handler)
	}()

    select {
    case err := <-errch:
        return err
    case <-ctx.Done():
        return srv.listener.Close()
    }
}

func Listen(handler http.Handler, opts *config.AppOptions) (*Server, error) {
	var (
		srv Server
		err error
	)

	srv.handler = handler
	srv.listener, err = net.Listen("tcp", opts.Addr())
	if err != nil {
		return nil, fmt.Errorf("failed to listen on port %d: %w", opts.Port, err)
	}

	switch opts.Mode {
	case config.HTTP:
		srv.serve = http.Serve
	case config.FCGI:
		srv.serve = fcgi.Serve
	default:
		return nil, config.ErrInvalidMode
	}

	return &srv, nil
}
