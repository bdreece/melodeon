package session

import (
	"github.com/gorilla/sessions"
	echoSession "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

const name string = "melodeon-session"

type Values struct {
	Username string `json:"username" mapstructure:"username"`
	Image    string `json:"image" mapstructure:"image"`
	Room     string `json:"room" mapstructure:"room"`
}

func Get(c echo.Context) (*sessions.Session, error) {
	return echoSession.Get(name, c)
}

func Decode(sess *sessions.Session) (*Values, error) {
	values := new(Values)
	if err := mapstructure.Decode(sess.Values, values); err != nil {
		return nil, err
	}

	return values, nil
}

func Encode(values Values, sess *sessions.Session) error {
	return mapstructure.Decode(values, &sess.Values)
}
