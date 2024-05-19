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

const (
	copyrightNotice string = "melodeon - Copyright (C) 2024 Brian Reece"
	licenseNotice   string = `
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.`
)

const (
	defaultPort       int    = 3000
	defaultConfigPath string = "/etc/melodeon/config.yml"
)

var args struct {
	Port       int
	ConfigPath string
}

var asRoute = dig.Group("routes")

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), copyrightNotice)
		_, _ = fmt.Fprintln(flag.CommandLine.Output())
		flag.PrintDefaults()
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), licenseNotice)
	}

	flag.IntVar(&args.Port, "p", defaultPort, "port")
	flag.StringVar(&args.ConfigPath, "c", defaultConfigPath, "config path")
	flag.Parse()
}

func main() {
	defer shutdown()

	cfg, err := config.Parse(args.ConfigPath)
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

	if err := app.Launch(ctx, args.Port); err != nil {
		panic(err)
	}
}

func shutdown() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "unexpected panic occurred: %v\n", r)
		os.Exit(1)
	}
}
