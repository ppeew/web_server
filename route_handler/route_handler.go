package route_handler

import (
	"net/http"
	"web_server/http_context"
)

type Routable interface {
	Route(method string, pattern string, handlerFunc func(c http_context.HttpContext))
}

// 路由分发处理
type RouteHandler interface {
	http.Handler
	Routable
}
