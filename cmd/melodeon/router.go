package main

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"github.com/bdreece/melodeon/pkg/router"
	"github.com/bdreece/melodeon/pkg/view"
)

var (
	asRoute = []dig.ProvideOption{
		dig.As(new(router.Route)),
		dig.Group("routes"),
	}

	asMiddleware = []dig.ProvideOption{
		dig.As(new(router.Middleware)),
		dig.Group("middlewares"),
	}
)

func createRouter(p struct {
	dig.In

	Routes      []router.Route      `group:"routes"`
	Middlewares []router.Middleware `group:"middlewares"`
	Renderer    echo.Renderer
	Validator   echo.Validator
	Logger      *slog.Logger
	Options     *view.Options
}) *echo.Echo {
	return router.New(&router.Options{
		Routes:      p.Routes,
		Middlewares: p.Middlewares,
		Renderer:    p.Renderer,
		Logger:      p.Logger,
		Options:     *p.Options,
	})
}
