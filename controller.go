package main

import "net/http"

type Controller struct {
	http.Handler
}

func NewController() *Controller {
	return &Controller{}
}
