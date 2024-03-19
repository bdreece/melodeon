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
	nonceKey    string = "nonce"
	tokenKey    string = "token"
	roomCodeKey string = "roomCode"

	DefaultCookie string = "__melodeon-session"
	NonceCookie   string = "__melodeon-nonce"
)

func init() {
	gob.Register(new(sessionToken))
	gob.Register(new(api.User))
}

type Session struct{ sessions.Session }

func (session Session) RoomCode() *string {
	if code, ok := session.Values[roomCodeKey].(string); ok && code != "" {
		return &code
	}

	return nil
}

func (session Session) Nonce() *string {
	if nonce, ok := session.Values[nonceKey].(string); ok && nonce != "" {
		return &nonce
	}

	return nil
}

func (session Session) User() *api.User {
	if user, ok := session.Values[userKey].(*api.User); ok && user != nil {
		return user
	}

	return nil
}

func (session Session) Token() *api.Token {
	if t, ok := session.Values[tokenKey].(*sessionToken); ok && t != nil {
		exp, _ := time.Parse(time.RFC3339, t.ExpiresIn)
		return &api.Token{
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
			TokenType:    t.TokenType,
			Scope:        t.Scope,
			ExpiresIn:    exp,
		}
	}

	return nil
}

func (session *Session) SetRoomCode(code string) { session.Values[roomCodeKey] = code }

func (session *Session) SetNonce(nonce string) { session.Values[nonceKey] = nonce }

func (session *Session) SetUser(user *api.User) { session.Values[userKey] = user }

func (session *Session) SetToken(token *api.Token) {
	if token == nil {
		session.Values[tokenKey] = nil
	} else {
		session.Values[tokenKey] = &sessionToken{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			TokenType:    token.TokenType,
			Scope:        token.Scope,
			ExpiresIn:    token.ExpiresIn.Format(time.RFC3339),
		}
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
