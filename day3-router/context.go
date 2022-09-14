package router

import "net/http"

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	// rquest info
	Path   string
	Method string
	Params map[string]string

	//response info
	StatusCode int
}

func (c *Context) Params(key string) string {
	value, _ := c.Params[key]
	return value
}
