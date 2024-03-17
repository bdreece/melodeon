package spotify

import (
	"crypto/rand"
	"encoding/base64"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/router/route"
)

var (
	authorizeEndpoint *url.URL
	authorizeScope = strings.Join([]string{
		"playlist-read-private",
		"playlist-read-collaborative",
		"streaming",
		"user-library-read",
		"user-modify-playback-state",
		"user-read-currently-playing",
		"user-read-playback-state",
	}, " ")
)

func init() {
	authorizeEndpoint, _ = url.Parse(
        "https://accounts.spotify.com/authorize?response_type=code")
}

type Authorize struct {
	route.Route

	log      *slog.Logger
	opts     *Options
	endpoint *url.URL
}

func (route *Authorize) Get(c echo.Context) error {
	endpoint := *route.endpoint
	state := make([]byte, 16)
	if _, err := rand.Read(state); err != nil {
		return logger.Error(route.log, "failed to generate state", err)
	}

	qs := endpoint.Query()
	qs.Set("state", base64.StdEncoding.EncodeToString(state))
	endpoint.RawQuery = qs.Encode()

	route.log.Info("redirecting to spotify authorization grant",
		slog.String("qs", endpoint.RawQuery))

	return c.Redirect(http.StatusFound, endpoint.String())
}

func NewAuthorize(opts *Options, log *slog.Logger) *Authorize {
	endpoint := *authorizeEndpoint
	qs := endpoint.Query()
	qs.Set("scope", authorizeScope)
	qs.Set("client_id", opts.ClientID)
	qs.Set("redirect_uri", opts.RedirectURI)
	endpoint.RawQuery = qs.Encode()

	return &Authorize{
		Route: route.New("/login"),

		log:      logger.For[Authorize](log),
		opts:     opts,
		endpoint: &endpoint,
	}
}

var (
    _ route.Route = (*Authorize)(nil)
    _ route.Get = (*Authorize)(nil)
)
