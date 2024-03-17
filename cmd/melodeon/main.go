package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/bdreece/melodeon/cmd/melodeon/app"
	"github.com/bdreece/melodeon/internal/guest"
	"github.com/bdreece/melodeon/internal/home"
	"github.com/bdreece/melodeon/internal/host"
	"github.com/bdreece/melodeon/pkg/contract"
	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify"
	"github.com/bdreece/melodeon/pkg/view"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

func main() {
	defer recoverer()

    cfgpath := flag.String("c", "configs/melodeon.json", "config path")
    flag.Parse()

    err := app.NewBuilder(*cfgpath).
		With(logger.New).
        With(session.NewStore).
        With(contract.NewValidator, dig.As(new(echo.Validator))).
		With(view.NewRenderer, dig.As(new(echo.Renderer))).
        With(view.NewNonceMiddleware, asMiddleware...).
		With(spotify.NewTokenClient).
		With(spotify.NewAuthorize, asRoute...).
		With(spotify.NewCallback, asRoute...).
		With(home.Default, asRoute...).
		With(host.DefaultPlayer, asRoute...).
		With(host.NewQueue, asRoute...).
		With(host.NewQueueItem, asRoute...).
		With(host.NewWizard, asRoute...).
		With(guest.NewRoom, asRoute...).
		With(createRouter, dig.As(new(http.Handler))).
		With(app.Listen).
		Build().
        Launch(context.Background())

    if err != nil {
        panic(err)
    }
}

func recoverer() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "unexpected panic occurred: %v\n", r)
		os.Exit(1)
	}
}
