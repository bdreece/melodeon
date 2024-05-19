package page

import (
	"net/http"

	"github.com/bdreece/yodel"
	"github.com/labstack/echo/v4"
)

const (
	homePath     string = "/"
	homeTemplate string = "home.gotmpl"
)

var Home = yodel.Route{
	Method: http.MethodGet,
	Path:   homePath,
	Handler: staticHandler{
		Template: homeTemplate,
		Data:     echo.Map{},
	},
}
