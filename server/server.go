package server

import "web_server/route_handler"

type Server interface {
	route_handler.Routable

	Start(address string) error
}
