package gee

import (
	"fmt"
	"net/http"
)

//HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

//Engine implement the interface of ServeHttp
