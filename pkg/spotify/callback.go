package spotify

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/bdreece/melodeon/pkg/spotify/api"
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
	var (
		users *UserClient
		token *api.Token
		user  *api.User
	)

	ctx := c.Request().Context()
	query, err := route.parseQuery(c)
	if err != nil {
		_ = logger.Error(route.log, "failed to parse query params", err)
		goto done
	}

	route.log.Info("handling callback request...",
		slog.String("state", query.state),
		slog.String("code", query.code))

	token, err = route.client.Exchange(ctx, query.code)
	if err != nil {
		_ = logger.Error(route.log, "failed to exchange code", err)
		goto done
	}

	route.log.Debug("exchanged authorization code",
		slog.String("type", token.TokenType),
		slog.String("scope", token.Scope),
		slog.String("expires_in", token.ExpiresIn.String()))

	users = &UserClient{
		client: client{
			Token:       *token,
			TokenClient: route.client,
		},
	}

	user, err = users.GetCurrentUser(ctx)
	if err != nil {
		_ = logger.Error(route.log, "failed to retrieve user profile", err)
		goto done
	}

	route.log.Debug("retrieved user profile",
		slog.String("display_name", user.DisplayName))

	if err = route.createSession(c, token, user); err != nil {
		_ = logger.Error(route.log, "failed to create session", err)
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

func (route *Callback) createSession(
	c echo.Context,
	token *api.Token,
	user *api.User,
) error {
	session, err := route.store.New(c, session.Name)
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
	client *TokenClient,
	store *session.Store,
	log *slog.Logger,
) *Callback {
	return &Callback{
		Route:  callbackRoute,
		client: client,
		store:  store,
		log:    logger.For[Callback](log),
	}
}
