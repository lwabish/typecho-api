package routes

import "github.com/gin-gonic/gin"

const (
	RootGroup = "root"
)

var (
	//routes "group": path,method,handler ...
	routes = map[RouteGroup][]RouteInfo{}
)

type RouteGroup string

type RouteInfo struct {
	Path       string
	Method     string
	Middleware []gin.HandlerFunc
	Handler    gin.HandlerFunc
}

func AddRoutes(group RouteGroup, ris []RouteInfo) {
	routes[group] = append(routes[group], ris...)
}

// RegisterRoutes register gin routes
func RegisterRoutes(router *gin.Engine) {
	for group, rs := range routes {
		if group == RootGroup {
			for _, route := range rs {
				router.Handle(route.Method, route.Path, append(route.Middleware, route.Handler)...)
			}
			continue
		}
		g := router.Group(string(group))
		for _, route := range rs {
			g.Handle(route.Method, route.Path, append(route.Middleware, route.Handler)...)
		}
	}
}
