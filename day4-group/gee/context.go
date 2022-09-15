package gee

import (
	"net/http"
)

type Context struct {
	//origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	//request info
	Path   string
	Method string
	Params map[string]string
	//response info
	StatusCode int
	//middleware
	handlers []HandlerFunc
	index    int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Path:   req.URL.Path,
		Method: req.Method,
		Req:    req,
		Writer: w,
		index:  -1,
	}
}

//在中间件中调用Next()方法时，控制权交给了下一个中间件
//c.Next()表示等待执行其他的中间件或者用户的handler
func (c *Context) Next() {
	c.index++ //index用于记录当前执行到第几个中间件
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}
