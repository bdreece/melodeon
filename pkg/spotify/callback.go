package spotify

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify/api"
	"github.com/bdreece/melodeon/pkg/store"
)

var (
	callbackRoute           = router.NewRoute("/callback")
	ErrMissingCallbackParam = errors.New("missing code or state param")
)

type Callback struct {
	router.Route

	handler  *TokenHandler
	sessions *session.Store
	store    *store.Store
	log      *slog.Logger
}

func (route *Callback) Get(c echo.Context) error {
	const redirect string = "/"
	var (
		users *UserClient
		token *api.Token
		user  *api.User
	)

	defer func() {
		_ = c.Redirect(http.StatusFound, redirect)
	}()

	ctx := c.Request().Context()
	query, err := route.parseQuery(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	route.log.Info("handling callback request...",
		slog.String("state", query.state),
		slog.String("code", query.code))

	token, err = route.handler.ExchangeCode(ctx, query.code)
	if err != nil {
		return echo.NewHTTPError(http.StatusFailedDependency, err.Error())
	}

	route.log.Debug("exchanged authorization code",
		slog.String("type", token.TokenType),
		slog.String("scope", token.Scope),
		slog.String("expires_in", token.ExpiresIn.String()))

	users = NewUserClient(token, route.handler)
	user, err = users.GetCurrentUser(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusFailedDependency, err.Error())
	}

	route.log.Debug("retrieved user profile",
		slog.String("display_name", user.DisplayName))

	if err = route.createSession(c, token, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (route *Callback) parseQuery(c echo.Context) (*struct {
	code  string
	state string
}, error) {
	qs := c.Request().URL.Query()
	code := qs.Get("code")
	state := qs.Get("state")
	if state == "" || code == "" {
		return nil, ErrMissingCallbackParam
	}

	return &struct {
		code  string
		state string
	}{
		code:  code,
		state: state,
	}, nil
}

func (route *Callback) createSession(
	c echo.Context,
	token *api.Token,
	user *api.User,
) error {
	session, err := route.sessions.New(c, session.DefaultCookie)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}

	session.SetUser(user)
	session.SetToken(token)
	if err = session.Save(c); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}

func NewCallback(
	handler *TokenHandler,
	sessions *session.Store,
	store *store.Store,
	log *slog.Logger,
) *Callback {
	return &Callback{
		Route:    callbackRoute,
		handler:  handler,
		sessions: sessions,
		store:    store,
		log:      logger.For[Callback](log),
	}
}
