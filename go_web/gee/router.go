package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}
}
