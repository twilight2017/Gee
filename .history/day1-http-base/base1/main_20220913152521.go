package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandlerFunc("/", indexHandler)
	http.HandlerFunc("/hello", helloHandler)
}
