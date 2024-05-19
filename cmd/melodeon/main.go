package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/bdreece/digdug"
	"github.com/bdreece/melodeon/internal/api"
	"github.com/bdreece/melodeon/internal/db"
	"github.com/bdreece/melodeon/internal/page"
	"github.com/bdreece/melodeon/internal/renderer"
	"github.com/bdreece/melodeon/internal/trace"
	"github.com/bdreece/melodeon/internal/validator"
	"github.com/bdreece/melodeon/internal/web"
	"github.com/bdreece/melodeon/pkg/config"
	"github.com/bdreece/melodeon/pkg/rooms"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

const defaultConfigPath string = "configs/melodeon.yml"

var configPath = flag.String("c", defaultConfigPath, "config path")

var (
	asRoute = dig.Group("routes")
)

func main() {
	defer shutdown()
	flag.Parse()

	cfg, err := config.Parse(*configPath)
	if err != nil {
		panic(err)
	}

	app := digdug.MustResolve[*web.App](
		digdug.New().
			Provide(digdug.Supply(&cfg.DB)).
			Provide(digdug.Supply(&cfg.Logging)).
			Provide(digdug.Supply(&cfg.Rooms)).
			Provide(digdug.Supply(&cfg.RootConfig)).
			Provide(digdug.Supply(&cfg.Sessions)).
			Provide(digdug.Supply(&cfg.Sessions.Cookie)).
			Provide(digdug.Supply(&cfg.Spotify)).
			Provide(digdug.Supply(&cfg.Web)).
			Provide(digdug.Supply(&cfg.Web.Config)).
			Provide(trace.NewLogger).
			Provide(db.New).
			Provide(session.NewStore, dig.As(new(sessions.Store))).
			Provide(rooms.NewStore).
			Provide(spotify.NewAuthManager).
			Provide(digdug.Supply(page.Home), asRoute).
			Provide(api.NewLoginRoute, asRoute).
			Provide(api.NewAuthorizeRoute, asRoute).
			Provide(renderer.New, dig.As(new(echo.Renderer))).
			Provide(digdug.Supply(validator.Default), dig.As(new(echo.Validator))).
			Provide(web.NewRouter, dig.As(new(http.Handler))).
			Provide(web.NewServer).
			Provide(web.NewApp).
			Container,
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.Launch(ctx); err != nil {
		panic(err)
	}
}

func shutdown() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "unexpected panic occurred: %v\n", r)
		os.Exit(1)
	}
}
