package core

import (
	"fmt"
	"net/http"
)

type Route struct {
	Req http.Request
}

func (route *Route) IsMethod(method string) bool {
	if route.Req.Method == method {
		return true
	} else {
		return false
	}
}

func (route *Route) IsPath(path string) bool {
	if route.Req.URL.Path == path {
		return true
	} else {
		return false
	}
}

func (route *Route) Get(path string, action func()) {
	if route.IsPath(path) {
		if route.IsMethod("GET") {
			action()
		} else {
			fmt.Println("Неверный метод, необходимо передать GET")
		}
	} else {
		fmt.Println("404 Not Found")
	}
}

func (route *Route) Post(path string, action func()) {
	if route.IsPath(path) {
		if route.IsMethod("Post") {
			action()
		} else {
			fmt.Println("Неверный метод, необходимо передать Post")
		}
	} else {
		fmt.Println("404 Not Found")
	}
}
