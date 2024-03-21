package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/bdreece/melodeon/cmd/melodeon/app"
	"github.com/bdreece/melodeon/internal/csp"
	"github.com/bdreece/melodeon/internal/page"
	"github.com/bdreece/melodeon/internal/page/guest"
	"github.com/bdreece/melodeon/internal/page/host"
	"github.com/bdreece/melodeon/internal/renderer"
	"github.com/bdreece/melodeon/internal/validator"
	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify"
	"github.com/bdreece/melodeon/pkg/store"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

func main() {
	defer recoverer()
	cfgpath := flag.String("c", "configs/melodeon.json", "config path")
	flag.Parse()

	program := app.NewBuilder(*cfgpath).
		With(logger.New).
		With(store.New).
		With(session.NewStore).
		With(spotify.NewTokenHandler).
		With(renderer.New, dig.As(new(echo.Renderer))).
		With(validator.Default, dig.As(new(echo.Validator))).
		With(csp.New, asMiddleware...).
		With(spotify.NewAuthorize, asRoute...).
		With(spotify.NewCallback, asRoute...).
		With(page.DefaultHome, asRoute...).
		With(page.NewLogout, asRoute...).
		With(host.DefaultPlayer, asRoute...).
		With(host.NewQueue, asRoute...).
		With(host.NewQueueItem, asRoute...).
		With(host.NewWizard, asRoute...).
		With(guest.NewRoom, asRoute...).
		With(createRouter, dig.As(new(http.Handler))).
		With(app.Listen).
		Build()

	if err := program.Launch(context.Background()); err != nil {
		panic(err)
	}
}

func recoverer() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "unexpected panic occurred: %v\n", r)
		os.Exit(1)
	}
}
