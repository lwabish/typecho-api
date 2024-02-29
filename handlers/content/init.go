package content

import (
	"github.com/lwabish/typecho-api/utils/routes"
	"net/http"
)

var (
	group routes.RouteGroup = "contents"
)

func init() {
	routes.AddRoutes(group, []routes.RouteInfo{
		{
			Path:    "",
			Method:  http.MethodPut,
			Handler: Hdl.Publish,
		},
	})
}
