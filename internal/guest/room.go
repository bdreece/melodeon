package guest

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/view"
	"github.com/bdreece/melodeon/pkg/router/route"
)

const roomTemplate string = "guest-room.gotmpl"
var roomRoute = route.New("/guest/room")

type Room struct {
	route.Route

	log *slog.Logger
}

// Get implements route.Get.
func (*Room) Get(c echo.Context) error {
    return c.Render(http.StatusOK, roomTemplate, view.Model{})
}

func NewRoom(log *slog.Logger) *Room {
	return &Room{roomRoute, logger.For[Room](log)}
}

var (
	_ route.Route = (*Room)(nil)
	_ route.Get   = (*Room)(nil)
)
