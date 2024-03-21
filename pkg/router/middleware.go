package router

import "github.com/labstack/echo/v4"

type Middleware interface {
	Invoke(next echo.HandlerFunc) echo.HandlerFunc
}

type MiddlewareFunc echo.MiddlewareFunc

func (f MiddlewareFunc) Invoke(next echo.HandlerFunc) echo.HandlerFunc {
	return f(next)
}
