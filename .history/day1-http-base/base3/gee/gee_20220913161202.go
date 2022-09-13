package gee

import (
	"fmt"
	"net/http"
)

//HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

//Engine implement the interface of ServeHttp
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of the gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine)POST(pattern string, handler HandlerFunc){
	engine.addRoute("POST", pattern, handler）
}
