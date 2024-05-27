package gorouter

import (
	"net/http"
)

type Router struct {
	c *Controller
	http.HandlerFunc
}

func NewRouter(c *Controller) *Router {
	return &Router{c: c}
}

func (rr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	switch method {
	case http.MethodGet:
		h, ok := rr.c.get[path]
		if !ok {
			var v Vars
			if h, v, ok = findPath(r, rr.c.get); ok {
				h.hFunc(toCtx(w, r, v))
				return
			}
			w.WriteHeader(http.StatusNotFound)
			return
		}
		h.hFunc(toCtx(w, r))
		break
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
