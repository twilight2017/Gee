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
