package gorouter

import (
	"fmt"
	"net/http"
	"strings"
)

func toCtx(w http.ResponseWriter, r *http.Request) *Ctx {
	return &Ctx{
		Response: w,
		Request:  r,
	}
}

func getVars(path string) map[int]string {
	resultParams := make(map[int]string)
	if path[0] == '/' {
		path = path[1:]
	}

	s := strings.Split(path, "/")
	counter := 0
	for i := range s {
		if strings.Contains(s[i], "{") {
			fmt.Println(s[i])
			param := fmt.Sprintf("%s", string(s[i][1:len(s[i])-1]))
			counter++
			resultParams[counter] = param
		}
	}
	return resultParams
}
