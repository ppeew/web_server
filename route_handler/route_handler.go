package route_handler

import (
	"net/http"
	"web_server/http_context"
)

// 路由分发处理
type RouteHandler interface {
	http.Handler
	Route(method string, pattern string, handlerFunc func(c http_context.HttpContext))
}
