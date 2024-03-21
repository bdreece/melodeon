package host

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/router"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/view"
)

const queueItemTemplate = "host-queue-item.gotmpl"

var queueItemRoute = router.NewRoute("/host/queue/{i}")

type QueueItem struct {
	router.Route

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
	_ router.Route       = (*QueueItem)(nil)
	_ router.GetRoute    = (*QueueItem)(nil)
	_ router.DeleteRoute = (*QueueItem)(nil)
	_ router.PutRoute    = (*QueueItem)(nil)
)
