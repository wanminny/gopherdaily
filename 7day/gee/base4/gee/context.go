package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request

	Path   string
	Method string

	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) PostForm(key string) string {
	return c.R.FormValue(key)
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.W.Header().Set("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.W.Header().Set("Content-Type", "text/html")
	c.Status(code)
	c.W.Write([]byte(html))

}

func (c *Context) SetHeader(key string, value string) {
	c.W.Header().Set(key, value)

}
