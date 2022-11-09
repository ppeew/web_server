package server

import (
	"net/http"
	"web_server/http_context"
	"web_server/route_handler"
)

var _ Server = NewDefaultServer("test")

type defaultHttpServer struct {
	Name string
	//使用的路由分发器
	routeHandler route_handler.RouteHandler
}

func (d *defaultHttpServer) Route(method string, pattern string, handlerFunc func(c http_context.HttpContext)) {
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	//hc := http_context.NewDefaultHttpContext(writer, request)
	//	//handlerFunc(hc)
	//	d.RouteHandler.Route(method, pattern, handlerFunc)
	//})

	//注册路由
	d.routeHandler.Route(method, pattern, handlerFunc)
}

func (d *defaultHttpServer) Start(address string) error {
	return http.ListenAndServe(address, route_handler.NewRouteHandlerBasedOnMap())
}

func NewDefaultServer(name string) Server {
	return &defaultHttpServer{
		Name:         name,
		routeHandler: route_handler.NewRouteHandlerBasedOnMap(),
	}
}
