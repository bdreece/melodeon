package api

import (
	"net/http"

	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/yodel"
	"github.com/labstack/echo/v4"
)

const logoutPath = "/logout"

func NewLogoutRoute() yodel.Route {
	return yodel.Route{
		Method: http.MethodGet,
		Path:   logoutPath,
		Handler: yodel.HandlerFunc(func(c echo.Context) error {
			sess, err := session.Get(c)
			if err != nil {
				return err
			}

			for k := range sess.Values {
				delete(sess.Values, k)
			}

			if err = sess.Save(c.Request(), c.Response()); err != nil {
				return err
			}

			return c.Redirect(http.StatusFound, "/")
		}),
	}
}
