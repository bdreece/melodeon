package router

import (
	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

var baseMiddlewares = []echo.MiddlewareFunc{
	middleware.Recover(),
    middleware.BodyLimit("2M"),
    middleware.Gzip(),
    middleware.Secure(),
    middleware.CSRF(),
}

func New(opts *Options) *echo.Echo {
	e := echo.New()
	e.Renderer = opts.Renderer
    for _, mw := range baseMiddlewares {
        e.Use(mw)
    }

	e.Use(slogecho.New(opts.Logger))
    e.Static(opts.StaticPrefix, opts.StaticDirectory)
    e.Static(opts.AppPrefix, opts.AppDirectory)

    for _, mw := range opts.Middlewares {
        e.Use(mw.Invoke)
    }

	for _, r := range opts.Routes {
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
