package web

import (
	"github.com/bdreece/yodel"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"go.uber.org/dig"

	"github.com/bdreece/melodeon/internal/trace"
	"github.com/bdreece/melodeon/pkg/config"
)

type RouterParams struct {
	dig.In
	Routes []yodel.Route `group:"routes"`

	Renderer     echo.Renderer
	Validator    echo.Validator
	SessionStore sessions.Store
	Log          *trace.Logger
	Config       *config.WebConfig
}

func NewRouter(p RouterParams) yodel.Router {
	router := yodel.New()
	router.Renderer = p.Renderer
	router.Validator = p.Validator
	router.Echo.Use(
		middleware.Recover(),
		middleware.Gzip(),
		middleware.BodyLimit("4M"),
		middleware.Secure(),
		session.Middleware(p.SessionStore),
		middleware.Static(p.Config.StaticDirectory),
		middleware.Static(p.Config.AssetsDirectory),
		slogecho.New(&p.Log.Logger),
	)

	for _, route := range p.Routes {
		router.Add(route)
	}

	return router
}
