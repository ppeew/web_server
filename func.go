package main

import (
	"web_server/data"
	"web_server/http_context"
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
