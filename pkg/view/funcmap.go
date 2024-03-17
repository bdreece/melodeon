package view

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

func funcMap(c echo.Context) template.FuncMap {
    return template.FuncMap{
        "nonce": func() *string {
            if nonce, ok := c.Get(nonceKey).(string); ok {
                return &nonce
            }

            return nil
        },
    }
}
