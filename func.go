package main

import (
	"io"
	"os"
	"web_server/data"
	"web_server/http_context"
)

func SignUpGet(c http_context.HttpContext) {
	//html返回
	//data:=
	//c.WriteJson(http.StatusOK,)
	file, err := os.OpenFile("./login.html", os.O_RDONLY, 0666)
	if err != nil {
		// TODO
	}
	data, err := io.ReadAll(file)
	if err != nil {
		// TODO
	}
	c.GetWriter().Header().Set("Content-Type", "text/html")
	c.GetWriter().Header().Set("charset", "UTF-8")
	c.OkJson(data)
}

func SignUpPost(c http_context.HttpContext) {
	req := data.SignUpReq{}
	err := c.ReadJson(&req)
	if err != nil {
		println(err.Error(), "读Body错误")
		c.BadRequest(nil)
		return
	}
	c.OkJson(req)
}
