package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandlerFunc("/", indexHandler)
	http.HandlerFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
