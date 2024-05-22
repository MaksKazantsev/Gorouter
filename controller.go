package gorouter

import "net/http"

type Controller struct {
	http.Handler
	h   Handler
	get map[string]HandlerStruct
}

func NewController() *Controller {
	return &Controller{
		get: map[string]HandlerStruct{},
	}
}

type Handler interface {
	GET(path string, fn HandlerFunc)
}

type HandlerStruct struct {
	hFunc HandlerFunc
}
type HandlerFunc func(*Ctx)
type Ctx struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (c *Controller) GET(path string, fn HandlerFunc) {
	c.get[path] = HandlerStruct{
		hFunc: fn,
	}
}
