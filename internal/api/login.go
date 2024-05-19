package api

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/bdreece/melodeon/pkg/spotify"
	"github.com/bdreece/yodel"
	"github.com/labstack/echo/v4"
)

const loginPath = "/login"

var (
	loginScopes = strings.Join([]string{}, " ")
	loginUrl, _ = url.Parse("https://accounts.spotify.com/authorize")
)

func NewLoginRoute(cfg *spotify.Config) yodel.Route {
	return yodel.Route{
		Method: http.MethodGet,
		Path:   loginPath,
		Handler: yodel.HandlerFunc(func(c echo.Context) error {
			var buffer bytes.Buffer
			if _, err := io.CopyN(&buffer, rand.Reader, 16); err != nil {
				return err
			}

			state := base64.URLEncoding.EncodeToString(buffer.Bytes())
			qs := url.Values{
				"response_type": {"code"},
				"client_id":     {cfg.Credentials.ClientID},
				"scope":         {loginScopes},
				"redirect_uri":  {cfg.RedirectURI},
				"state":         {state},
			}

			uri := *loginUrl
			uri.RawQuery = qs.Encode()

			return c.Redirect(http.StatusFound, uri.String())
		}),
	}
}
