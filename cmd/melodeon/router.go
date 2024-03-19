package main

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"github.com/bdreece/melodeon/pkg/router"
	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/view"
)

var (
	asRoute = []dig.ProvideOption{
		dig.As(new(route.Route)),
		dig.Group("routes"),
	}

	asMiddleware = []dig.ProvideOption{
		dig.As(new(route.Middleware)),
		dig.Group("middlewares"),
	}
)

func createRouter(p struct {
	dig.In

	Routes      []route.Route      `group:"routes"`
	Middlewares []route.Middleware `group:"middlewares"`
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
