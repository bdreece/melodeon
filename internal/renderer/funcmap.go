package renderer

import (
	"html/template"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/session"
)

func funcMap(c echo.Context) template.FuncMap {
	return template.FuncMap{
		"session": func() *session.Values {
			sess, err := session.Get(c)
			if err != nil {
				c.Logger().Error(err.Error())
				return nil
			}

			values, err := session.Decode(sess)
			if err != nil {
				c.Logger().Error(err.Error())
				return nil
			}

			if values.Room == "" && values.Image == "" && values.Username == "" {
				return nil
			}

			return values
		},

		"unescapeJS":   unescapeJS,
		"unescapeHTML": unescapeHTML,
	}
}

func unescapeJS(s string) template.JS {
	return template.JS(s)
}

func unescapeHTML(s string) template.HTML {
	return template.HTML(s)
}
