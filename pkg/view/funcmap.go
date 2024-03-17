package view

import (
	"html/template"

	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify"
	"github.com/labstack/echo/v4"
)

func FuncMap(c echo.Context, session *session.Session) template.FuncMap {
    return template.FuncMap{
        "user": func() *spotify.User {
            return session.Values["user"].(*spotify.User)
        },
    }
}
