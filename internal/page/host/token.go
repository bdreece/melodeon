package host

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/session"
)

var tokenRoute = route.New("/host/token")

type Token struct {
	route.Route

	sessions *session.Store
	log      *slog.Logger
}

func (t *Token) Get(c echo.Context) error {
	type response struct {
		Token string `json:"token"`
	}

	s, err := t.sessions.Get(c, session.DefaultCookie)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token := s.Token()
	if token == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "no token in session")
	}

	return c.JSON(http.StatusOK, response{
		Token: token.AccessToken,
	})
}

func NewToken(sessions *session.Store, log *slog.Logger) *Token {
	return &Token{tokenRoute, sessions, logger.For[Token](log)}
}

var (
	_ route.Route = (*Token)(nil)
	_ route.Get   = (*Token)(nil)
)
