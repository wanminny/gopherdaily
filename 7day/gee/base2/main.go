package main

import (
	"fmt"
	"log"
	"net/http"
)

type engine struct {
}

func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/index.html":
		//w.Write([]byte("<h1>Hello, world!</h1>"))
		fmt.Fprintf(w, "url path:%q", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "%v:%v\n", k, v)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<h1>404 Not Found</h1>"))
	}
}

func main() {

	log.Fatalln(http.ListenAndServe(":8080", &engine{}))

}
