package page

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/view"
)

const homeTemplate string = "home.gotmpl"

var homeRoute = route.New("/")
var defaultHome = Home{homeRoute}

type Home struct{ route.Route }

// Get implements Get.
func (Home) Get(c echo.Context) error {
	return c.Render(http.StatusOK, homeTemplate, view.Model{})
}

func DefaultHome() *Home { return &defaultHome }

var (
	_ route.Route = (*Home)(nil)
	_ route.Get   = (*Home)(nil)
)
