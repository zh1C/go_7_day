package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct {}

func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found:%s\n", req.URL.Path)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}

