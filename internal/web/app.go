package web

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/bdreece/melodeon/internal/trace"
)

type App struct {
	server  Server
	handler http.Handler
	log     *trace.Logger
}

func (app *App) Launch(ctx context.Context, port int) error {
	defer app.log.Info("goodbye!")

	addr := fmt.Sprintf(":%d", port)
	lst, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to open tcp listener on address %q: %v", addr, err)
	}

	errch := make(chan error, 1)
	go func() {
		defer close(errch)
		app.log.Info("launching application", slog.String("addr", addr))
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
) *App {
	return &App{
		server:  server,
		handler: handler,
		log:     log,
	}
}
