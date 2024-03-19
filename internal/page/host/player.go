package host

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/view"
)

const playerTemplate string = "host-player.gotmpl"
var playerRoute = route.New("/host/player")
var defaultPlayer = Player{playerRoute}

type Player struct{ route.Route }

// Get implements route.Get.
func (Player) Get(c echo.Context) error {
	return c.Render(http.StatusOK, playerTemplate, view.Model{})
}

func DefaultPlayer() *Player { return &defaultPlayer }

var (
	_ route.Route = (*Player)(nil)
	_ route.Get   = (*Player)(nil)
)
