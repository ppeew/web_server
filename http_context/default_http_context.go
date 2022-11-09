package http_context

import (
	"encoding/json"
	"io"
	"net/http"
)

var _ HttpContext = NewDefaultHttpContext(nil, nil)

type DefaultHttpContext struct {
	W http.ResponseWriter
	R *http.Request
}

func (d *DefaultHttpContext) GetWriter() http.ResponseWriter {
	return d.W
}

func (d *DefaultHttpContext) GetReader() *http.Request {
	return d.R
}

func (d *DefaultHttpContext) ReadJson(data interface{}) error {
	body, err := io.ReadAll(d.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func (d *DefaultHttpContext) WriteJson(state int, data interface{}) error {
	resp := d.W
	resp.WriteHeader(state)
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = resp.Write(b)
	return err
}

func (d *DefaultHttpContext) OkJson(data interface{}) error {
	return d.WriteJson(http.StatusOK, data)
}

func (d *DefaultHttpContext) SystemErrJson(data interface{}) error {
	return d.WriteJson(http.StatusInternalServerError, data)
}
func (d *DefaultHttpContext) BadRequest(data interface{}) error {
	return d.WriteJson(http.StatusBadRequest, data)
}

func NewDefaultHttpContext(w http.ResponseWriter, r *http.Request) *DefaultHttpContext {
	return &DefaultHttpContext{
		W: w,
		R: r,
	}
}
