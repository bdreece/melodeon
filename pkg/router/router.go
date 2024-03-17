package router

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"

	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/view"
)

func New(
	routes []route.Route,
	renderer echo.Renderer,
	log *slog.Logger,
    opts *view.Options,
) *echo.Echo {
	e := echo.New()
	e.Renderer = renderer
	e.Use(middleware.Recover())
	e.Use(slogecho.New(log))
    e.Static(opts.StaticPrefix, opts.StaticDirectory)
    e.Static(opts.AppPrefix, opts.AppDirectory)

	for _, r := range routes {
		if h, ok := r.(route.Get); ok {
			e.GET(r.Pattern(), h.Get)
		}
		if h, ok := r.(route.Post); ok {
			e.POST(r.Pattern(), h.Post)
		}
		if h, ok := r.(route.Put); ok {
			e.PUT(r.Pattern(), h.Put)
		}
		if h, ok := r.(route.Patch); ok {
			e.PATCH(r.Pattern(), h.Patch)
		}
		if h, ok := r.(route.Delete); ok {
			e.DELETE(r.Pattern(), h.Delete)
		}
	}

	return e
}
