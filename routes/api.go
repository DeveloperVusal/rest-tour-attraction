package routes

import (
	"attrtour/core"
	"net/http"
)

type Api struct{}

func (api *Api) Run(wri http.ResponseWriter, req *http.Request) {
	router := core.HandleRouter{
		Req: *req,
		Wri: wri,
	}
	route := core.Route{
		Routes:       map[string]string{},
		RoutesAction: map[string]func(){},
	}

	// User paths
	route.Get("/api/user/get", func() {
		wri.Write([]byte("Get user by Id=578"))
	})

	route.Post("/api/user/add", func() {
		wri.Write([]byte("Add new user"))
	})

	route.Post("/api/place/add", func() {
		wri.Write([]byte("Add new place"))
	})

	// Run Router Handle
	router.TunnelControl(&route.Routes, &route.RoutesAction)
}
