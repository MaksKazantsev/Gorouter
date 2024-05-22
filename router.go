package main

import "net/http"

type Router struct {
	c *Controller
}

func NewRouter(c *Controller) *Router {
	return &Router{c: c}
}

func (rr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
