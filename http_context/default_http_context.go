package http_context

import (
	"encoding/json"
	"io"
	"net/http"
)

var _ HttpContext = NewDefaultHttpContext(nil, nil)

type defaultHttpContext struct {
	W http.ResponseWriter
	R *http.Request
}

func (d *defaultHttpContext) ReadJson(data interface{}) error {
	body, err := io.ReadAll(d.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func (d *defaultHttpContext) WriteJson(state int, data interface{}) error {
	resp := d.W
	resp.WriteHeader(state)
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = resp.Write(b)
	return err
}

func (d *defaultHttpContext) OkJson(data interface{}) error {
	return d.WriteJson(http.StatusOK, data)
}

func (d *defaultHttpContext) SystemErrJson(data interface{}) error {
	return d.WriteJson(http.StatusInternalServerError, data)
}
func (d *defaultHttpContext) BadRequest(data interface{}) error {
	return d.WriteJson(http.StatusBadRequest, data)
}

func NewDefaultHttpContext(w http.ResponseWriter, r *http.Request) *defaultHttpContext {
	return &defaultHttpContext{
		W: w,
		R: r,
	}
}
