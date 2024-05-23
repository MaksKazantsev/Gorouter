package gorouter

import "net/http"

type Router struct {
	c *Controller
	http.HandlerFunc
}

func NewRouter(c *Controller) *Router {
	return &Router{c: c}
}

func (rr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	url := r.URL.Path

	switch method {
	case http.MethodGet:
		handler, ok := rr.c.get[url]
		if !ok {
			findPath(r, rr.c.get)
			w.WriteHeader(http.StatusNotFound)
		}
		handler.hFunc(toCtx(w, r))
		break
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
