package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync/atomic"
	"web_server/filter_builder"
	"web_server/http_context"
	"web_server/route_handler"
)

var _ Server = NewDefaultServer("test")

type DefaultHttpServer struct {
	Name string
	//使用的路由分发器
	routeHandler route_handler.RouteHandler
	root         filter_builder.Filter

	//用于关闭的
	closing       uint32
	requestCount  int64
	zeroRequestCh chan struct{}
}

func (d *DefaultHttpServer) Shutdown(ctx context.Context) error {
	return d.RejectNewRequestAndWaiting(ctx)
	//这里不需要释放资源
}

func (d *DefaultHttpServer) Route(method string, pattern string, handlerFunc func(c http_context.HttpContext)) {
	//注册路由
	d.routeHandler.Route(method, pattern, handlerFunc)
}

func (d *DefaultHttpServer) Start(address string) error {
	return http.ListenAndServe(address, d)
}

func (d *DefaultHttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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
	return &DefaultHttpServer{
		Name:          name,
		routeHandler:  routeHandler,
		root:          root,
		zeroRequestCh: make(chan struct{}, 1),
	}
}

func (d *DefaultHttpServer) RejectNewRequestAndWaiting(ctx context.Context) error {
	atomic.AddUint32(&d.closing, 1)
	if atomic.LoadInt64(&d.requestCount) == 0 {
		return nil
	}

	select {
	case <-ctx.Done():
		return errors.New("RejectNewRequestAndWaiting problem : timeout")
	case <-d.zeroRequestCh:
		fmt.Println("all request finished")
	}
	return nil
}

// 用于过滤用的,当服务器关闭，拒绝新连接
func (g *DefaultHttpServer) ShutdownFilterBuilder(f filter_builder.Filter) filter_builder.Filter {
	return func(c http_context.HttpContext) {
		if atomic.LoadUint32(&g.closing) >= 1 {
			c.WriteJson(http.StatusServiceUnavailable, nil)
			return
		}
		atomic.AddInt64(&g.requestCount, 1)
		f(c)
		atomic.AddInt64(&g.requestCount, -1)
		if atomic.LoadInt64(&g.requestCount) == 0 {
			g.zeroRequestCh <- struct{}{}
		}
	}
}
