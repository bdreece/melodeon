package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

var baseMiddlewares = []echo.MiddlewareFunc{
	middleware.Recover(),
	middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)),
	middleware.BodyLimit("2M"),
	middleware.Gzip(),
	middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge: 0,
	}),
	middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "cookie:__melodeon-csrf",
		CookieName:     "__melodeon-csrf",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}),
	middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}),
}

func New(opts *Options) *echo.Echo {
	e := echo.New()
	e.Renderer = opts.Renderer
	e.Validator = opts.Validator
	e.HTTPErrorHandler = handleError

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
		if h, ok := r.(GetRoute); ok {
			e.GET(r.Pattern(), h.Get)
		}
		if h, ok := r.(PostRoute); ok {
			e.POST(r.Pattern(), h.Post)
		}
		if h, ok := r.(PutRoute); ok {
			e.PUT(r.Pattern(), h.Put)
		}
		if h, ok := r.(PatchRoute); ok {
			e.PATCH(r.Pattern(), h.Patch)
		}
		if h, ok := r.(DeleteRoute); ok {
			e.DELETE(r.Pattern(), h.Delete)
		}
	}

	return e
}
