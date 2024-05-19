package trace

import (
	"log/slog"

	"github.com/gofrs/uuid"
)

type Trace struct {
	slog.Logger

	ID uuid.UUID
}

type Span struct {
	Trace

	Parent uuid.UUID
}

func (log *Trace) Span() *Span {
	id, _ := uuid.NewV4()

	return &Span{
		Parent: log.ID,
		Trace: Trace{
			ID: id,
			Logger: *log.With(slog.Group("span",
				slog.String("id", id.String()),
				slog.String("parent", log.ID.String()),
			)),
		},
	}
}
