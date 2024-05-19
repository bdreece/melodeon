package api

import (
	"log/slog"
	"net/http"

	"github.com/bdreece/yodel"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	"github.com/bdreece/melodeon/internal/trace"
	"github.com/bdreece/melodeon/pkg/rooms"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify"
)

type authorizeQuery struct {
	State string `query:"state" validate:"required"`
	Code  string `query:"code" validate:"required_without=Error"`
	Error string `query:"error"`
}

type authorizeHandler struct {
	auth   *spotify.AuthManager
	config *session.CookieConfig
	rooms  *rooms.Store
	log    *trace.Logger
}

func (h *authorizeHandler) Handle(c echo.Context) error {
	trace := h.log.Trace()
	query := new(authorizeQuery)
	if err := c.Bind(query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if err := c.Validate(query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if query.Error != "" {
		return echo.NewHTTPError(http.StatusFailedDependency, query.Error)
	}

	trace.Info("authorizing request...", slog.String("code", query.Code))
	authorizer := h.auth.User(query.Code)
	ctx := c.Request().Context()
	token, err := authorizer.Authorize(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	client := spotify.UserClient{
		Client: spotify.Client{
			Authorizer: authorizer,
		},
	}

	trace.Info("fetching profile...")
	user, err := client.Profile(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	trace.Info("creating and storing room...")
	key, _ := uuid.NewV4()
	room := rooms.NewRoom(user.DisplayName, token)
	if err := h.rooms.Put(key, *room, trace); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	trace.Info("creating session...")
	sess, err := session.Get(c)
	if err != nil {
		return err
	}

	sess.Options = h.config.SessionOptions()
	values := session.Values{
		Username: user.DisplayName,
		Room:     key.String(),
	}

	if len(user.Images) > 0 {
		values.Image = user.Images[0].Url
	}

	if err := session.Encode(values, sess); err != nil {
		return err
	}

	trace.Info("saving session...")
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/")
}

type AuthorizeRouteOptions struct {
	dig.In

	AuthManager *spotify.AuthManager
	Config      *session.CookieConfig
	Rooms       *rooms.Store
	Log         *trace.Logger
}

func NewAuthorizeRoute(opts AuthorizeRouteOptions) yodel.Route {
	return yodel.Route{
		Method: http.MethodGet,
		Path:   "/authorize",
		Handler: &authorizeHandler{
			auth:   opts.AuthManager,
			config: opts.Config,
			rooms:  opts.Rooms,
			log:    opts.Log,
		},
	}
}
