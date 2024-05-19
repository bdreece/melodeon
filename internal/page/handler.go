package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
type loader interface {
	Load(echo.Context) (echo.Map, error)
}

type loaderFunc func(echo.Context) (echo.Map, error)
*/

type staticHandler struct {
	Template string
	Data     echo.Map
}

func (h staticHandler) Handle(c echo.Context) error {
	return c.Render(http.StatusOK, h.Template, h.Data)
}

/*
type dynamicHandler struct {
	Template string
	Loader   loader
}

func (h dynamicHandler) Handle(c echo.Context) error {
	data, err := h.Loader.Load(c)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, h.Template, data)
}
*/
