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
	GET(path string, fn HandlerFunc, mw ...Middleware)
	POST(path string, fn HandlerFunc, mw ...Middleware)
	DELETE(path string, fn HandlerFunc, mw ...Middleware)
	PUT(path string, fn HandlerFunc, mw ...Middleware)
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
func (c *Controller) POST(path string, fn HandlerFunc, mw ...Middleware) {
	for _, middleware := range mw {
		fn = middleware(fn)
	}

	elems, vars := handlePath(path)

	c.get[path] = HandlerStruct{
		hFunc: fn,
		elems: elems,
		vars:  vars,
	}
}

// POST is HTTP post method
func (g *group) POST(path string, fn HandlerFunc, mw ...Middleware) {
	for _, middleware := range mw {
		fn = middleware(fn)
	}

	elems, vars := handlePath(path)

	g.c.get[g.base+path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// DELETE is HTTP delete method
func (g *group) DELETE(path string, fn HandlerFunc, mw ...Middleware) {
	for _, middleware := range mw {
		fn = middleware(fn)
	}

	elems, vars := handlePath(path)

	g.c.get[g.base+path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// DELETE is HTTP delete method
func (c *Controller) DELETE(path string, fn HandlerFunc, mw ...Middleware) {
	for _, middleware := range mw {
		fn = middleware(fn)
	}

	elems, vars := handlePath(path)

	c.get[path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// GET is HTTP get method
func (g *group) GET(path string, fn HandlerFunc, mw ...Middleware) {
	for _, middleware := range mw {
		fn = middleware(fn)
	}

	elems, vars := handlePath(path)

	g.c.get[g.base+path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// GET is HTTP get method
func (c *Controller) GET(path string, fn HandlerFunc, mw ...Middleware) {
	for _, middleware := range mw {
		fn = middleware(fn)
	}

	elems, vars := handlePath(path)

	c.get[path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// PUT is HTTP put method
func (c *Controller) PUT(path string, fn HandlerFunc, mw ...Middleware) {
	for _, middleware := range mw {
		fn = middleware(fn)
	}

	elems, vars := handlePath(path)

	c.get[path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// PUT is HTTP put method
func (g *group) PUT(path string, fn HandlerFunc, mw ...Middleware) {
	for _, middleware := range mw {
		fn = middleware(fn)
	}

	elems, vars := handlePath(path)

	g.c.get[g.base+path] = HandlerStruct{
		hFunc: fn,
		vars:  vars,
		elems: elems,
	}
}

// HandlerStruct represents all parts, from which Handler consist of.
type HandlerStruct struct {
	hFunc HandlerFunc

	vars  map[int]string
	elems []string
}

// HandlerFunc is a func that handles http request
type HandlerFunc func(*Ctx)

// Ctx represents response, request and variables and passes into HandlerFunc
type Ctx struct {
	Response http.ResponseWriter
	Request  *http.Request
	Vars     Vars
}

// Vars represents variables in the path, which can be declared by {variable}. This variables can be obtained from this map by their key
type Vars map[string]string
