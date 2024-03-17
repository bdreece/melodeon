package host

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/view"
)

const queueItemTemplate = "host-queue-item.gotmpl"
var queueItemRoute = route.New("/host/queue/{i}")

type QueueItem struct {
	route.Route

	store *session.Store
	log   *slog.Logger
}

// Put implements route.Put.
func (*QueueItem) Put(echo.Context) error {
	panic("unimplemented")
}

// Delete implements route.Delete.
func (*QueueItem) Delete(echo.Context) error {
	panic("unimplemented")
}

// Get implements route.Get.
func (QueueItem) Get(c echo.Context) error {
    return c.Render(http.StatusOK, queueItemTemplate, view.Model{})
}

func NewQueueItem(store *session.Store, log *slog.Logger) *QueueItem {
	return &QueueItem{queueItemRoute, store, log}
}

var (
	_ route.Route  = (*QueueItem)(nil)
	_ route.Get    = (*QueueItem)(nil)
	_ route.Delete = (*QueueItem)(nil)
	_ route.Put    = (*QueueItem)(nil)
)
