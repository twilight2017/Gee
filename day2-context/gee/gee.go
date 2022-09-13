package gee

import (
	"net/http"
)

//Handler defines the request handler used by gee
type HandlerFunc func(*Context)

//Engine implement the interface of ServeHTTp
type Engine struct {
	router *router
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

//Get defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//Run defines a method to start a http serve
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(":9999", engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req, http.Request){
	c := newContext(w, req)
	engine.router.handle(c)
}
