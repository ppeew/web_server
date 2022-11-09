package route_handler

import (
	"fmt"
	"net/http"
	"web_server/http_context"
)

// 检查是否实现接口
var _ RouteHandler = NewRouteHandlerBasedOnMap()

type routeHandlerBasedOnMap struct {
	m map[string]func(c http_context.HttpContext)
}

func (r *routeHandlerBasedOnMap) Route(method string, pattern string, handlerFunc func(c http_context.HttpContext)) {
	//添加路由
	key := r.key(method, pattern)
	fmt.Printf("添加%s\n", key)
	r.m[key] = handlerFunc
}

func (r *routeHandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//找到路由既接受，否则拒绝
	key := r.key(request.Method, request.URL.Path)
	fmt.Printf("查找%s\n", key)
	if f, ok := r.m[key]; ok {
		//找到 做路由处理
		c := http_context.NewDefaultHttpContext(writer, request)
		f(c)
	} else {
		//找不到
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not any route match"))
		return
	}
}

func (r *routeHandlerBasedOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}

func NewRouteHandlerBasedOnMap() *routeHandlerBasedOnMap {
	return &routeHandlerBasedOnMap{}
}
