package session

import (
	"encoding/gob"
	"time"

	"github.com/bdreece/melodeon/pkg/spotify/api"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const (
	userKey     string = "user"
	tokenKey    string = "token"
	roomCodeKey string = "roomCode"

    Name string = "__melodeon-session"
)

func init() {
    gob.Register(new(sessionToken))
    gob.Register(new(api.User))
}

type Session struct{ sessions.Session }

func (session Session) RoomCode() *string {
    if code, ok := session.Values[roomCodeKey].(string); ok {
        return &code
    }

    return nil
}


func (session Session) User() *api.User  {
    if user, ok := session.Values[userKey].(*api.User); ok {
        return user
    }

    return nil
}

func (session Session) Token() *api.Token {
	t := session.Values[tokenKey].(*sessionToken)
	exp, _ := time.Parse(time.RFC3339, t.ExpiresIn)
	return &api.Token{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		TokenType:    t.TokenType,
		Scope:        t.Scope,
		ExpiresIn:    exp,
	}
}

func (session *Session) SetUser(user *api.User)  { session.Values[userKey] = user }

func (session *Session) SetRoomCode(code string) { session.Values[roomCodeKey] = code }

func (session *Session) SetToken(token *api.Token) {
    session.Values[tokenKey] = &sessionToken{
        AccessToken: token.AccessToken,
        RefreshToken: token.RefreshToken,
        TokenType: token.TokenType,
        Scope: token.Scope,
        ExpiresIn: token.ExpiresIn.Format(time.RFC3339),
    }
}

func (session *Session) Save(c echo.Context) error {
    return session.Session.Save(c.Request(), c.Response().Writer)
}

type sessionToken struct {
	AccessToken  string
	RefreshToken string
	Scope        string
	TokenType    string
	ExpiresIn    string
}

