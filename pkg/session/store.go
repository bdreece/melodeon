package session

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type Store struct {
    sessions.Store
}

func NewStore(opts *Options) *Store {
    store := sessions.NewCookieStore(opts.SecretKey)
    store.Options = &opts.Options
    return &Store{store}
}

func (store *Store) Get(c echo.Context, name string) (*Session, error) {
    s, err := store.Store.Get(c.Request(), name)
    if err != nil {
        return nil, err
    }

    return &Session{*s}, nil
}

func (store *Store) New(c echo.Context, name string) (*Session, error) {
    s, err := store.Store.New(c.Request(), name)
    if err != nil {
        return nil, err
    }

    return &Session{*s}, nil
}

func (store *Store) Save(c echo.Context, session *Session) error {
    return store.Store.Save(c.Request(), c.Response().Writer, &session.Session)
}
