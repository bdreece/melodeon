package host

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/view"
)

const queueTemplate string = "host-queue.gotmpl"

var queueRoute = route.New("/host/queue")

type Queue struct {
	route.Route

	store *session.Store
	log   *slog.Logger
}

// Get implements route.Get.
func (Queue) Get(c echo.Context) error {
	return c.Render(http.StatusOK, queueTemplate, view.Model{})
}

// Post implements route.Post.
func (Queue) Post(echo.Context) error {
	panic("unimplemented")
}

// Patch implements route.Patch.
func (Queue) Patch(echo.Context) error {
	panic("unimplemented")
}

// Delete implements route.Delete.
func (Queue) Delete(echo.Context) error {
	panic("unimplemented")
}

func NewQueue(store *session.Store, log *slog.Logger) *Queue {
	return &Queue{queueRoute, store, logger.For[Queue](log)}
}

var (
	_ route.Route  = (*Queue)(nil)
	_ route.Get    = (*Queue)(nil)
	_ route.Patch  = (*Queue)(nil)
	_ route.Post   = (*Queue)(nil)
	_ route.Delete = (*Queue)(nil)
)
