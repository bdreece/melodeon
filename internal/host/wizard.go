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

const wizardTemplate = "host-wizard.gotmpl"

var wizardRoute = route.New("/host/wizard")

type Wizard struct {
	route.Route

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
	_ route.Route = (*Wizard)(nil)
	_ route.Get   = (*Wizard)(nil)
	_ route.Post  = (*Wizard)(nil)
)
