package main

import "base4/gee"

func main() {
	// 请在此处添加代
	eng := gee.NewEngine()
	eng.GET("/", func(ctx *gee.Context) {
		ctx.HTML(200, "<h1>Hello Gee</h1>")
	})
	eng.GET("/hello", func(ctx *gee.Context) {
		ctx.String(200, "hello,%s, %s", ctx.Query("username"), ctx.PostForm("password"))
	})
	eng.POST("/hello1", func(ctx *gee.Context) {
		ctx.JSON(200, "hello json")
	})
	eng.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(200, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password")})
	})
	eng.Start(":8080")
}
