package server

import "web_server/http_context"

type Server interface {
	//命中路由时处理函数hanlerFunc
	Route(method string, pattern string, handlerFunc func(c http_context.HttpContext))
	Start(address string) error
}
