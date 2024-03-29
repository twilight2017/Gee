package main

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		//start time
		t := time.Now()
		// if a server error occured
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())  //global middleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) //v2 group middlegroup
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			//expect /hello/geektutu
			c.String(http.StatusOk, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.run(":9999")
}
