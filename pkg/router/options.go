package router

import (
	"log/slog"

	"github.com/bdreece/melodeon/pkg/view"
	"github.com/labstack/echo/v4"
)

type Options struct {
	view.Options

	Routes      []Route
	Middlewares []Middleware
	Renderer    echo.Renderer
	Validator   echo.Validator
	Logger      *slog.Logger
}
