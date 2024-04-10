package main

import "base5/gee"

func main() {

	eng := gee.NewEngine()

	//eng.GET("/", func(c *gee.Context) {
	//	c.JSON(200, gee.H{
	//		"message": "Hello Gee",
	//	})
	//})
	//eng.GET("/hello", func(c *gee.Context) {
	//	c.String(200, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//})

	eng.GET("/hello/:name", func(c *gee.Context) {
		c.String(200, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	//eng.GET("/assets/*file", func(c *gee.Context) {
	//	c.JSON(200, gee.H{
	//		"file": c.Param("file"),
	//	})
	//})
	eng.Run(":8080")
}
