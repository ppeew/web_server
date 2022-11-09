package main

import (
	"web_server/data"
	"web_server/http_context"
	"web_server/server"
)

func SignUp(c http_context.HttpContext) {
	req := data.SignUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		c.BadRequest(nil)
		return
	}
	c.OkJson(nil)
}

func main() {
	svr := server.NewDefaultServer("defaultServer")
	svr.Route("GET", "/user_login", SignUp)
	err := svr.Start(":8080")
	if err != nil {
		panic(err)
	}
}
