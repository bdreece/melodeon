package session

import (
	"log/slog"
	"os"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type Store struct {
	sessions.Store

	log *slog.Logger
}

func (store *Store) Get(c echo.Context, name string) (*Session, error) {
	s, err := store.Store.Get(c.Request(), name)
	if err != nil {
		return nil, logger.Error(store.log, "failed to get session", err)
	}

	return &Session{*s}, nil
}

func (store *Store) New(c echo.Context, name string) (*Session, error) {
	s, err := store.Store.New(c.Request(), name)
	if err != nil {
		return nil, logger.Error(store.log, "failed to create session", err)
	}

	return &Session{*s}, nil
}

func (store *Store) Save(c echo.Context, session *Session) error {
	return store.Store.Save(c.Request(), c.Response().Writer, &session.Session)
}

func NewStore(log *slog.Logger, opts *Options) *Store {
	store := sessions.NewFilesystemStore(os.TempDir(), opts.SecretKey)
	store.Options = &opts.Cookies
	return &Store{store, logger.For[Store](log)}
}
