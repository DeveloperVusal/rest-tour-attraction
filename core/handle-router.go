package core

import (
	"net/http"
)

type HandleRouter struct {
	Req http.Request
	Wri http.ResponseWriter
}

func (hr *HandleRouter) TunnelControl(routes *map[string]string, actions *map[string]func()) {
	isNotFound := true

	for path, method := range *routes {
		if hr.IsValidPath(path) {
			isNotFound = false

			if hr.IsMethod(method) {
				(*actions)[path]()
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

// func (hr *HandleRouter) IsRoute(path string, method string) bool {
// 	if hr.Routes[path] == method {
// 		return true
// 	} else {
// 		return false
// 	}
// }
