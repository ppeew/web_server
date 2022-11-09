package http_context

import "net/http"

type HttpContext interface {
	ReadJson(data interface{}) error
	WriteJson(state int, data interface{}) error
	OkJson(data interface{}) error
	SystemErrJson(data interface{}) error
	BadRequest(data interface{}) error
	GetWriter() http.ResponseWriter
	GetReader() *http.Request
}
