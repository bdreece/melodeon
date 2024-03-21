package guest

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router"
	"github.com/bdreece/melodeon/pkg/view"
)

const roomTemplate string = "guest-room.gotmpl"

var roomRoute = router.NewRoute("/guest/room")

type Room struct {
	router.Route

	log *slog.Logger
}

// Get implements route.Get.
func (route *Room) Get(c echo.Context) error {
	req := new(RoomRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	route.log.Info("request", slog.String("code", req.Code))

	return c.Render(http.StatusOK, roomTemplate, view.Model{"Req": req})
}

func NewRoom(log *slog.Logger) *Room {
	return &Room{roomRoute, logger.For[Room](log)}
}

var (
	_ router.Route    = (*Room)(nil)
	_ router.GetRoute = (*Room)(nil)
)

type RoomRequest struct {
	Code string `param:"code"`
}
