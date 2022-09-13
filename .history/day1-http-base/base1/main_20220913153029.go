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

func indexHandler(w http.ResponseWriter, req *http.Request{
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
})

func helloHandler(w http.ResponseWriter, req *hhtp.Request){
	for, k,v:= range Header{
		fmt.FPrintf(w, "Header[%q] = %q\n", k, v)
	}
}
