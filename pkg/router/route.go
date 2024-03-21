package router

import "github.com/labstack/echo/v4"

type (
	route struct {
		pattern string
	}

	Route interface {
		Pattern() string
	}

	GetRoute interface {
		Get(echo.Context) error
	}

	PostRoute interface {
		Post(echo.Context) error
	}

	PutRoute interface {
		Put(echo.Context) error
	}

	PatchRoute interface {
		Patch(echo.Context) error
	}

	DeleteRoute interface {
		Delete(echo.Context) error
	}

	MiddlewareRoute interface {
		Middlewares() []Middleware
	}
)

func (r route) Pattern() string { return r.pattern }

func NewRoute(pattern string) Route {
	return route{pattern}
}
