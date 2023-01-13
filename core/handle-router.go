package core

import (
	"fmt"
	"net/http"
)

type HandleRouter struct {
	Req http.Request
	Wri http.ResponseWriter
}

func (hr *HandleRouter) TunnelControl(routes *map[string][]interface{}) {
	isNotFound := true

	for path, interf := range *routes {
		if hr.IsValidPath(path) {
			isNotFound = false
			method := interf[0].(string)

			if hr.IsMethod(method) {
				f_ptr, ok := interf[1].(func())

				if !ok {
					panic(fmt.Sprintf("Invalid object type: expected `func()`, turned out to be `%T`", interf[1]))
				}

				f_ptr()
				hr.Wri.WriteHeader(http.StatusOK)
			} else {
				hr.Wri.WriteHeader(http.StatusMethodNotAllowed)
				hr.Wri.Write([]byte("Method Not Allowed\n"))
				hr.Wri.Write([]byte(method + ", must be passed\n"))
			}
		}
	}

	if isNotFound {
		hr.Wri.WriteHeader(http.StatusNotFound)
		hr.Wri.Write([]byte("404 Not Found\n"))
	}
}

func (hr *HandleRouter) IsMethod(method string) bool {
	if hr.Req.Method == method {
		return true
	} else {
		return false
	}
}

func (hr *HandleRouter) IsValidPath(path string) bool {
	if hr.Req.URL.Path == path {
		return true
	} else {
		return false
	}
}
