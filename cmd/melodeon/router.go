package main

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"github.com/bdreece/melodeon/pkg/router"
	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/view"
)

var asRoute = []dig.ProvideOption{
	dig.As(new(route.Route)),
	dig.Group("routes"),
}

func newRouter(p struct {
	dig.In

	Routes   []route.Route `group:"routes"`
	Renderer echo.Renderer
	Log      *slog.Logger
	Options  *view.Options
}) *echo.Echo {
	return router.New(p.Routes, p.Renderer, p.Log, p.Options)
}
