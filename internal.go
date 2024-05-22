package gorouter

import "net/http"

func toCtx(w http.ResponseWriter, r *http.Request) *Ctx {
	return &Ctx{
		Response: w,
		Request:  r,
	}
}
