package server

import (
	"net/http"
	"web_server/filter_builder"
	"web_server/http_context"
	"web_server/route_handler"
)

var _ Server = NewDefaultServer("test")

type defaultHttpServer struct {
	Name string
	//使用的路由分发器
	routeHandler route_handler.RouteHandler
	root         filter_builder.Filter
}

func (d *defaultHttpServer) Route(method string, pattern string, handlerFunc func(c http_context.HttpContext)) {
	//注册路由
	d.routeHandler.Route(method, pattern, handlerFunc)
}

func (d *defaultHttpServer) Start(address string) error {
	return http.ListenAndServe(address, d)
}

func (d *defaultHttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := http_context.NewDefaultHttpContext(writer, request)
	//调用责任链
	d.root(c)
}

func NewDefaultServer(name string, builders ...filter_builder.FilterBuilder) Server {
	routeHandler := route_handler.NewRouteHandlerBasedOnMap()
	var root filter_builder.Filter = routeHandler.ServeHTTP
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}
	return &defaultHttpServer{
		Name:         name,
		routeHandler: routeHandler,
		root:         root,
	}
}
