package route

import "github.com/labstack/echo/v4"

type Get interface {
	Get(echo.Context) error
}

type Post interface {
	Post(echo.Context) error
}

type Put interface {
    Put(echo.Context) error
}

type Patch interface {
	Patch(echo.Context) error
}

type Delete interface {
	Delete(echo.Context) error
}
