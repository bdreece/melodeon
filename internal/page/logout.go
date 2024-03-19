package page

import (
	"net/http"

	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/labstack/echo/v4"
)

var logoutRoute = route.New("/logout")

type Logout struct {
	route.Route

	sessions *session.Store
}

func (route *Logout) Get(c echo.Context) error {
	defer func() {
        _ = c.Redirect(http.StatusFound, "/")
    }()

	cookie, err := c.Cookie(session.DefaultCookie)
	if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cookie.MaxAge = -1
	c.SetCookie(cookie)
    return nil
}

func NewLogout(sessions *session.Store) *Logout {
	return &Logout{logoutRoute, sessions}
}

var (
	_ route.Route = (*Logout)(nil)
	_ route.Get   = (*Logout)(nil)
)
