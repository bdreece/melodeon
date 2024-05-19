package api

import (
	"net/http"

	"github.com/bdreece/melodeon/pkg/rooms"
	"github.com/bdreece/yodel"
	"github.com/labstack/echo/v4"
)

const joinPath = "/join"

type joinHandler struct {
	store *rooms.Store
}

func (h *joinHandler) Handle(c echo.Context) error {
	return nil
}

func NewJoinRoute(store *rooms.Store) yodel.Route {
	return yodel.Route{
		Method: http.MethodPost,
		Path:   joinPath,
		Handler: &joinHandler{
			store: store,
		},
	}
}
