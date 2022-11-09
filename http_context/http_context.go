package http_context

type HttpContext interface {
	ReadJson(data interface{}) error
	WriteJson(state int, data interface{}) error
	OkJson(data interface{}) error
	SystemErrJson(data interface{}) error
	BadRequest(data interface{}) error
}
