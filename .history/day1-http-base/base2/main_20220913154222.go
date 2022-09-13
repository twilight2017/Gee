package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServerHttp(w http.ResponseWriter, req *http.Request)
