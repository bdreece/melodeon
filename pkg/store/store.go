package store

import (
	"context"
	"encoding/json"
	"io/fs"
	"log/slog"

	"github.com/bdreece/melodeon/pkg/logger"
	"github.com/bdreece/melodeon/pkg/spotify/api"
	"github.com/boltdb/bolt"
)

const (
	tokenBucketKey string = "tokens"
)

type Store struct {
	*bolt.DB

	log *slog.Logger
}

func (s *Store) Get(ctx context.Context, key string) (*api.Token, error) {
	type result struct {
		data *api.Token
		err  error
	}

	ch := make(chan result, 1)
	defer close(ch)
	go func() {
		tx, err := s.DB.Begin(false)
		if err != nil {
			ch <- result{err: err}
			return
		}

		token := new(api.Token)
		bucket := tx.Bucket([]byte(tokenBucketKey))
		if err := json.Unmarshal(bucket.Get([]byte(key)), &token); err != nil {
			ch <- result{err: err}
			return
		}

		ch <- result{data: token}
	}()

	select {
	case res := <-ch:
		return res.data, res.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *Store) Put(ctx context.Context, key string, token *api.Token) error {
	ch := make(chan error, 1)
	defer close(ch)
	go func() {
		tx, err := s.DB.Begin(true)
		if err != nil {
			ch <- err
			return
		}

		bucket := tx.Bucket([]byte(tokenBucketKey))
		value, err := json.Marshal(token)
		if err != nil {
			ch <- err
			return
		}

		if err = bucket.Put([]byte(key), value); err != nil {
			ch <- err
			return
		}

		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func New(log *slog.Logger, opts *Options) (*Store, error) {
	const perms fs.FileMode = 0o0600

	db, err := bolt.Open(opts.Path, perms, nil)
	if err != nil {
		return nil, err
	}

	return &Store{db, logger.For[Store](log)}, nil
}
