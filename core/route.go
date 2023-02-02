package core

type Route struct {
	Routes map[string][]interface{}
	Prefix string
}

func (route *Route) Get(path string, action func(map[string]interface{})) {
	method := "GET"

	route.Routes[route.Prefix+path] = []interface{}{method, action}
}

func (route *Route) Post(path string, action func(map[string]interface{})) {
	method := "POST"

	route.Routes[route.Prefix+path] = []interface{}{method, action}
}

func (route *Route) Put(path string, action func(map[string]interface{})) {
	method := "PUT"

	route.Routes[route.Prefix+path] = []interface{}{method, action}
}

func (route *Route) Patch(path string, action func(map[string]interface{})) {
	method := "PATCH"

	route.Routes[route.Prefix+path] = []interface{}{method, action}
}

func (route *Route) Delete(path string, action func(map[string]interface{})) {
	method := "DELETE"

	route.Routes[route.Prefix+path] = []interface{}{method, action}
}
