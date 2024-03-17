package router

import (
	"log/slog"

	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/view"
	"github.com/labstack/echo/v4"
)

type Options struct {
    view.Options

	Routes      []route.Route
	Middlewares []route.Middleware
	Renderer    echo.Renderer
	Logger      *slog.Logger
}
