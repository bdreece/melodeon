package route

import "github.com/labstack/echo/v4"

type Middleware interface {
    Invoke(next echo.HandlerFunc) echo.HandlerFunc
}
