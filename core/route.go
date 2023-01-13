package core

type Route struct {
	Routes map[string][]interface{}
}

func (route *Route) Get(path string, action func()) {
	method := "GET"

	route.Routes[path] = []interface{}{method, action}
}

func (route *Route) Post(path string, action func()) {
	method := "POST"

	route.Routes[path] = []interface{}{method, action}
}

func (route *Route) Put(path string, action func()) {
	method := "PUT"

	route.Routes[path] = []interface{}{method, action}
}

func (route *Route) Patch(path string, action func()) {
	method := "PATCH"

	route.Routes[path] = []interface{}{method, action}
}

func (route *Route) Delete(path string, action func()) {
	method := "DELETE"

	route.Routes[path] = []interface{}{method, action}
}
