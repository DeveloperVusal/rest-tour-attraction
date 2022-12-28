package core

type Route struct {
	Routes       map[string]string
	RoutesAction map[string]func()
}

func (route *Route) Get(path string, action func()) {
	method := "GET"

	route.Routes[path] = method
	route.RoutesAction[path] = action
}

func (route *Route) Post(path string, action func()) {
	method := "POST"

	route.Routes[path] = method
	route.RoutesAction[path] = action
}

func (route *Route) Put(path string, action func()) {
	method := "PUT"

	route.Routes[path] = method
	route.RoutesAction[path] = action
}

func (route *Route) Patch(path string, action func()) {
	method := "PATCH"

	route.Routes[path] = method
	route.RoutesAction[path] = action
}

func (route *Route) Delete(path string, action func()) {
	method := "DELETE"

	route.Routes[path] = method
	route.RoutesAction[path] = action
}
