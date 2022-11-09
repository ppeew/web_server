package server

import (
	"net/http"
	"web_server/route_handler"
)

type Server interface {
	route_handler.Routable
	http.Handler
	Start(address string) error
}
