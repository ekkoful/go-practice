package main

import (
	"ekkogee/ekkogee"
	"log"
	"net/http"
	"time"
)

func main() {

	r := ekkogee.New()

	r.GET("/", func(c *ekkogee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	r.GET("/panic", func(c *ekkogee.Context) {
		names := []string{"ekkoful"}
		c.String(http.StatusOK, names[100])
	})

	v1 := r.Group("/v1")
	v1.GET("/", func(c *ekkogee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ekkoGee</h1>")
	})

	v1.GET("/hello", func(c *ekkogee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *ekkogee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}
	r.Run(":9999")
}

func onlyForV2() ekkogee.HandlerFunc {
	return func(c *ekkogee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
