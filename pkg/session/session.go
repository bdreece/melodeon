package session

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

const (
	roomCodeKey     string = "roomCode"
	userNameKey     string = "userName"
	accessTokenKey  string = "accessToken"
	refreshTokenKey string = "refreshToken"
	expirationKey   string = "expiration"
)

type Session struct{ sessions.Session }

func (session Session) RoomCode() string { return session.Values[roomCodeKey].(string) }
func (session Session) UserName() string { return session.Values[userNameKey].(string) }
func (session Session) AccessToken() string { return session.Values[accessTokenKey].(string) }
func (session Session) RefreshToken() string { return session.Values[refreshTokenKey].(string) }
func (session Session) Expiration() time.Time {
    tstr := session.Values[expirationKey].(string)
    t, _ := time.Parse(time.RFC3339, tstr)
    return t
}

func (session *Session) SetRoomCode(code string) {
	session.Values[roomCodeKey] = code
}

func (session *Session) SetUserName(userName string) {
	session.Values[userNameKey] = userName
}

func (session *Session) SetAccessToken(accessToken string) {
	session.Values[accessTokenKey] = accessToken
}

func (session *Session) SetRefreshToken(refreshToken string) {
	session.Values[refreshTokenKey] = refreshToken
}

func (session *Session) SetExpiration(expiration time.Time) {
	session.Values[expirationKey] = expiration.Format(time.RFC3339)
}

func Get(store sessions.Store, r *http.Request, name string) (*Session, error) {
    s, err := store.Get(r, name)
    if err != nil {
        return nil, err
    }

    return &Session{*s}, nil
}
