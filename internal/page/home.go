package page

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/router"
	"github.com/bdreece/melodeon/pkg/view"
)

const homeTemplate string = "home.gotmpl"

var homeRoute = router.NewRoute("/")
var defaultHome = Home{homeRoute}

type Home struct{ router.Route }

// Get implements router.GetRoute.
func (Home) Get(c echo.Context) error {
	return c.Render(http.StatusOK, homeTemplate, view.Model{})
}

func DefaultHome() *Home { return &defaultHome }

var (
	_ router.Route    = (*Home)(nil)
	_ router.GetRoute = (*Home)(nil)
)
