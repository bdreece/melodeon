package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/bdreece/melodeon/pkg/logger"
)

type App struct {
	srv *Server
	log *slog.Logger
}

func (app *App) Launch(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		<-quit
		cancel()
	}()

	app.log.Info("launching server...",
		slog.String("addr", app.srv.Addr()))

	if err := app.srv.Serve(ctx); err != nil {
		return logger.Error(app.log, "server shutdown unexpectedly", err)
	}

	return nil
}

func New(srv *Server, log *slog.Logger) *App {
	return &App{srv, logger.For[App](log)}
}
