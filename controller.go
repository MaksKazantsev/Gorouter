package gorouter

import (
	"net/http"
)

type Controller struct {
	http.Handler
	h struct {
		Handler
		Group
	}
	get    map[string]HandlerStruct
	post   map[string]HandlerStruct
	delete map[string]HandlerStruct
	put    map[string]HandlerStruct
}

func NewController() *Controller {
	return &Controller{
		get:    make(map[string]HandlerStruct),
		post:   make(map[string]HandlerStruct),
		delete: make(map[string]HandlerStruct),
		put:    make(map[string]HandlerStruct),
	}
}

type Group interface {
	Group(base string) Handler
}

type Handler interface {
	GET(path string, fn HandlerFunc)
	POST(path string, fn HandlerFunc)
	DELETE(path string, fn HandlerFunc)
	PUT(path string, fn HandlerFunc)
}

// Group groups several handlers by one base string
func (c *Controller) Group(base string) Handler {
	return &group{
		c:    c,
		base: base,
	}
}

type group struct {
	c    *Controller
	base string
}

// POST is HTTP post method
func (c *Controller) POST(path string, fn HandlerFunc) {
	elems, vars := handlePath(path)

	c.get[path] = HandlerStruct{
		hFunc: fn,
		elems: elems,
		vars:  vars,
	}
}

// POST is HTTP post method
func (g *group) POST(path string, fn HandlerFunc) {
	elems, vars := handlePath(path)
	g.c.get[g.base+path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// DELETE is HTTP delete method
func (g *group) DELETE(path string, fn HandlerFunc) {
	elems, vars := handlePath(path)
	g.c.get[g.base+path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// DELETE is HTTP delete method
func (c *Controller) DELETE(path string, fn HandlerFunc) {
	elems, vars := handlePath(path)

	c.get[path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// GET is HTTP get method
func (g *group) GET(path string, fn HandlerFunc) {
	elems, vars := handlePath(path)

	g.c.get[g.base+path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// GET is HTTP get method
func (c *Controller) GET(path string, fn HandlerFunc) {
	elems, vars := handlePath(path)

	c.get[path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// PUT is HTTP put method
func (c *Controller) PUT(path string, fn HandlerFunc) {
	elems, vars := handlePath(path)

	c.get[path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// PUT is HTTP put method
func (g *group) PUT(path string, fn HandlerFunc) {
	elems, vars := handlePath(path)

	g.c.get[g.base+path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

type HandlerStruct struct {
	hFunc HandlerFunc

	vars  map[string]string
	elems []string
}
type HandlerFunc func(*Ctx)
type Ctx struct {
	Response   http.ResponseWriter
	Request    *http.Request
	Parameters Parameters
}
type Parameters map[string]string
