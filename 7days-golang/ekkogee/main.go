package main

import (
	"ekkogee/ekkogee"
	"net/http"
)

func main() {

	r := ekkogee.New()

	r.GET("/", func(c *ekkogee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ekkoGee</h1>")
	})

	r.GET("/hello", func(c *ekkogee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *ekkogee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *ekkogee.Context) {
		c.JSON(http.StatusOK, ekkogee.H{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(c *ekkogee.Context) {
		c.JSON(http.StatusOK, ekkogee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
