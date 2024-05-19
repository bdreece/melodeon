package web

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/bdreece/melodeon/internal/trace"
	"github.com/bdreece/melodeon/pkg/config"
)

type App struct {
	addr    string
	server  Server
	handler http.Handler
	log     *trace.Logger
}

func (app *App) Launch(ctx context.Context) error {
	defer app.log.Info("goodbye!")

	lst, err := net.Listen("tcp", app.addr)
	if err != nil {
		return fmt.Errorf("failed to open tcp listener on address %q: %v", app.addr, err)
	}

	errch := make(chan error, 1)
	go func() {
		defer close(errch)
		app.log.Info("launching application", slog.String("addr", app.addr))
		if err = app.server.Serve(lst, app.handler); err != nil && err != net.ErrClosed {
			errch <- err
		}
	}()

	select {
	case err := <-errch:
		return err
	case <-ctx.Done():
		break
	}

	app.log.Info("shutting down...")
	if err = lst.Close(); err != nil {
		return err
	}

	return nil
}

func NewApp(
	server Server,
	handler http.Handler,
	log *trace.Logger,
	cfg *config.RootConfig,
) *App {
	return &App{
		addr:    fmt.Sprintf(":%d", cfg.Port),
		server:  server,
		handler: handler,
		log:     log,
	}
}
