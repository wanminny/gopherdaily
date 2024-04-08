package main

import (
	"base3/gee"
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}

func main() {

	gee := gee.New()
	gee.GET("/", indexHandler) // 这里可以写代码
	gee.POST("/hello", helloHandler)
	gee.Start(":8080")
}
