package rooms

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/bdreece/melodeon/internal/trace"
	"github.com/boltdb/bolt"
	"github.com/gofrs/uuid"
)

var ErrNotFound = errors.New("room not found")

type Store struct {
	db     *bolt.DB
	config *Config
}

func (store *Store) Get(key uuid.UUID, log *trace.Trace) (*Room, error) {
	var span *trace.Span
	if log != nil {
		span = log.Span()
	}

	rx, err := store.db.Begin(false)
	if err != nil {
		return nil, fmt.Errorf("failed to begin boltdb transaction: %v", err)
	}

	if span != nil {
		span.Debug("opening rooms bucket", slog.String("name", store.config.BucketName))
	}

	bucket, err := store.ensureBucketCreated(
		rx, rx.Bucket([]byte(store.config.BucketName)), &span.Logger)

	if err != nil {
		return nil, err
	}

	if span != nil {
		span.Debug("querying rooms bucket", slog.String("key", key.String()))
	}

	data := bucket.Get(key.Bytes())
	if data == nil {
		return nil, ErrNotFound
	}

	if span != nil {
		span.Debug("unmarshaling room JSON", slog.String("data", string(data)))
	}

	room := new(Room)
	if err := json.Unmarshal(data, room); err != nil {
		return nil, fmt.Errorf("failed to unmarshal room entry for key %q: %v", key.String(), err)
	}

	return room, nil
}

func (store *Store) Put(key uuid.UUID, room Room, log *trace.Trace) error {
	var span *trace.Span
	if log != nil {
		span = log.Span()
	}

	return store.db.Update(func(tx *bolt.Tx) error {
		if span != nil {
			span.Debug("opening rooms bucket", slog.String("name", store.config.BucketName))
		}

		bucket, err := store.ensureBucketCreated(
			tx, tx.Bucket([]byte(store.config.BucketName)), &span.Logger)

		if err != nil {
			return err
		}

		if span != nil {
			span.Debug("marshaling room JSON", slog.Any("room", room))
		}

		data, err := json.Marshal(room)
		if err != nil {
			return fmt.Errorf("failed to marshal room entry: %v", err)
		}

		if span != nil {
			span.Debug("putting room data",
				slog.String("key", key.String()),
				slog.String("data", string(data)))
		}

		if err = bucket.Put(key.Bytes(), data); err != nil {
			return fmt.Errorf("failed to put room entry %q: %v", key.String(), err)
		}

		return nil
	})
}

func NewStore(db *bolt.DB, cfg *Config) *Store { return &Store{db, cfg} }

func (store Store) ensureBucketCreated(
	tx *bolt.Tx,
	bucket *bolt.Bucket,
	log *slog.Logger,
) (*bolt.Bucket, error) {
	if bucket != nil {
		return bucket, nil
	}

	if log != nil {
		log.Debug("creating rooms bucket", slog.String("name", store.config.BucketName))
	}

	bucket, err := tx.CreateBucket([]byte(store.config.BucketName))
	if err != nil {
		return nil, fmt.Errorf("failed to create bucket %q: %v", store.config.BucketName, err)
	}

	return bucket, nil
}
