package gorouter

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func toCtx(w http.ResponseWriter, r *http.Request, v ...Vars) *Ctx {
	c := &Ctx{
		Request:  r,
		Response: w,
	}

	if len(v) > 0 {
		c.Vars = v[0]
	}

	return c
}

// handlePath ha
func handlePath(path string) ([]string, map[int]string) {
	pathElems := make([]string, 0)
	paramsValue := make(map[int]string)
	if path[0] == '/' {
		path = path[1:]
	}

	s := strings.Split(path, "/")
	for i := range s {
		if strings.Contains(s[i], "{") && strings.Contains(s[i], "}") {
			param := fmt.Sprintf("%s", string(s[i][1:len(s[i])-1]))
			paramsValue[i+1] = param
		} else {
			pathElems = append(pathElems, s[i])
		}
	}
	return pathElems, paramsValue
}

// parseVars parses variables {var} and write them into map
func parseVars(u *url.URL, params map[int]string) map[string]string {
	if len(params) < 1 {
		return nil
	}

	path := u.Path

	if path[0] == '/' {
		path = path[1:]
	}

	s := strings.Split(path, "/")

	vars := make(map[string]string)

	for k, v := range params {
		vars[v] = s[k-1]
	}

	return vars
}

// findPath finds a path, which is mostly like the one that was provided in a handlers map.
func findPath(r *http.Request, h map[string]HandlerStruct) (HandlerStruct, Vars, bool) {
	path := r.URL.Path
	if path[0] == '/' {
		path = path[1:]
	}

	doneCh := make(chan struct{}, len(h))
	foundCh := make(chan struct {
		h HandlerStruct
		v Vars
	}, 1)
	defer func() {
		close(doneCh)
		close(foundCh)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, handler := range h {
		go func(ctx context.Context, h HandlerStruct) {
			if len(h.vars) < 1 {
				select {
				case <-ctx.Done():
				default:
					doneCh <- struct{}{}
				}
				return
			}

			for _, elements := range handler.elems {
				if !strings.Contains(path, elements) {
					select {
					case <-ctx.Done():
					default:
						doneCh <- struct{}{}
					}
					return
				}
			}
			vars := parseVars(&url.URL{Path: path}, h.vars)

			select {
			case <-ctx.Done():
			default:
				foundCh <- struct {
					h HandlerStruct
					v Vars
				}{h: h, v: vars}
			}
		}(ctx, handler)
	}
	doneCounter := 0
	for {
		select {
		case f := <-foundCh:
			cancel()
			return f.h, f.v, true
		case <-doneCh:
			doneCounter++
			if doneCounter == len(h) {
				cancel()
				return HandlerStruct{}, nil, false
			}
		}
	}
}
