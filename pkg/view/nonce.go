package view

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log/slog"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router/route"
	"github.com/labstack/echo/v4"
)

const nonceKey string = "nonce"

type NonceMiddleware struct {
	log *slog.Logger
}

// Invoke implements route.Middleware.
func (mw *NonceMiddleware) Invoke(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		buf := make([]byte, 16)
		if _, err := io.ReadFull(rand.Reader, buf); err != nil {
			panic(logger.Error(mw.log, "failed to generate nonce", err))
		}

		c.Set(nonceKey, base64.StdEncoding.EncodeToString(buf))
		return next(c)
	}
}

func NewNonceMiddleware(log *slog.Logger) *NonceMiddleware {
	return &NonceMiddleware{logger.For[NonceMiddleware](log)}
}

var _ route.Middleware = (*NonceMiddleware)(nil)
