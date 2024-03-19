package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/view"
)

const errorTemplate string = "error.gotmpl"

func handleError(err error, c echo.Context) {
    code := http.StatusInternalServerError
    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
    }

    c.Logger().Error(err)
    if err = c.Render(code, errorTemplate, view.Model{
        "Error": err.Error(),
    }); err != nil {
        panic(err)
    }
}
