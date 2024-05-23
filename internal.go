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

func saveVars(path string) ([]string, map[string]string) {
	paramsName := make([]string, 0)
	paramsValue := make(map[string]string)
	if path[0] == '/' {
		path = path[1:]
	}

	s := strings.Split(path, "/")
	for i := range s {
		if strings.Contains(s[i], "{") && strings.Contains(s[i], "}") {
			param := fmt.Sprintf("%s", string(s[i][1:len(s[i])-1]))
			paramsValue[param] = "0"
			paramsName = append(paramsName, param)
		}
	}
	return paramsName, paramsValue
}
