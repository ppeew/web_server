package main

import (
	"web_server/filter_builder"
	"web_server/server"
	"web_server/shutdown"
)

func main() {
	svr := server.NewDefaultServer("defaultServer", filter_builder.TimeFilterBuilder)
	//退出函数 在这里需要：1.关闭所有Server 2.取消用户请求 3.处理完所有的用户的请求 4.释放资源  否则.超时强制关闭
	go shutdown.WaitForShutdown(shutdown.BuildCloseServersHook(svr))
	svr.Route("GET", "/user_login", SignUp)
	err := svr.Start("localhost:8080")
	if err != nil {
		panic(err)
	}
}
