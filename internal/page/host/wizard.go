package host

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/view"
)

const wizardTemplate = "host-wizard.gotmpl"

var wizardRoute = router.NewRoute("/host/wizard")

type Wizard struct {
	router.Route

	store *session.Store
	log   *slog.Logger
}

// Get implements route.Get.
func (route *Wizard) Get(c echo.Context) error {
	return c.Render(http.StatusOK, wizardTemplate, view.Model{})
}

// Post implements route.Post.
func (route *Wizard) Post(c echo.Context) error {
	panic("unimplemented")
}

func NewWizard(store *session.Store, log *slog.Logger) *Wizard {
	return &Wizard{wizardRoute, store, logger.For[Wizard](log)}
}

var (
	_ router.Route     = (*Wizard)(nil)
	_ router.GetRoute  = (*Wizard)(nil)
	_ router.PostRoute = (*Wizard)(nil)
)
