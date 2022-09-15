package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})

	//index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.COntext) {
		names := []string{"dada"}
		c.String(http.StatusOK, names[100])
	})
	r.RUN(":9999")
}
