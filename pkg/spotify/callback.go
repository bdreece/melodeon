package spotify

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/session"
)

var (
	callbackRoute           = route.New("/callback")
	ErrMissingCallbackParam = errors.New("missing code or state param")
)

type Callback struct {
	route.Route

	client *TokenClient
	store  *session.Store
	log    *slog.Logger
}

func (route *Callback) Get(c echo.Context) error {
	const redirect string = "/"
	var res *Token

	query, err := route.parseQuery(c)
	if err != nil {
		_ = logger.Error(route.log, "failed to parse query params", err)
		goto done
	}

	route.log.Info("handling callback request...",
		slog.String("state", query.state),
		slog.String("code", query.code))

	res, err = route.client.Exchange(c.Request().Context(), query.code)
	if err != nil {
		_ = logger.Error(route.log, "failed to exchange code", err)
		goto done
	}

	route.log.Debug("exchanged authorization code",
		slog.String("type", res.TokenType),
		slog.String("scope", res.Scope),
		slog.Int("expires_in", res.ExpiresIn))

	if err = route.updateSession(c, res); err != nil {
		_ = logger.Error(route.log, "failed to update session", err)
		goto done
	}

done:
	return c.Redirect(http.StatusFound, redirect)
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

func (route *Callback) updateSession(c echo.Context, res *Token) error {
	expiresAt := time.Now().Add(time.Duration(res.ExpiresIn) * time.Second)
	session, err := route.store.Get(c, "__melodeon-host")
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}

	session.SetAccessToken(res.AccessToken)
	session.SetRefreshToken(res.RefreshToken)
	session.SetExpiration(expiresAt)

	err = route.store.Save(c, session)
	if err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}

func NewCallback(client *TokenClient, store *session.Store, log *slog.Logger) *Callback {
	return &Callback{callbackRoute, client, store, logger.For[Callback](log)}
}
