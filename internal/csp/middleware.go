package csp

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log/slog"
	"net/http"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/bdreece/melodeon/pkg/session"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	sessions *session.Store
	log      *slog.Logger
}

func (mw *Middleware) Invoke(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := mw.sessions.Get(c, session.NonceCookie)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		nonce := sess.Nonce()
		if nonce == nil {
			buf := make([]byte, 8)
			if _, err := io.ReadFull(rand.Reader, buf); err != nil {
				panic(logger.Error(mw.log, "failed to generate nonce", err))
			}

			v := hex.EncodeToString(buf)
			sess.SetNonce(v)
			if err = sess.Save(c); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			nonce = &v
		}

		policy := DefaultPolicy
		policy.ScriptSrc = append(policy.ScriptSrc, "'nonce-"+*nonce+"'")

		c.Response().Header().Add("Content-Security-Policy", policy.String())
		return next(c)
	}
}

func New(sessions *session.Store, log *slog.Logger) *Middleware {
	return &Middleware{sessions, log}
}

var _ route.Middleware = (*Middleware)(nil)
