package routes

import (
	"attrtour/core"
	"net/http"
)

type Web struct {
	IsFindRoute *bool
}

func (web *Web) Run(wri http.ResponseWriter, req *http.Request) {
	router := core.HandleRouter{
		Req: req,
		Wri: wri,
	}
	route := core.Route{
		Routes: map[string][]interface{}{},
	}

	// Render Vue
	route.Get("/", func(args map[string]interface{}) {
		view := core.View{}

		view.Render(wri, "/views/vue/dist/index.html")
	})

	// Run Router Handle
	router.TunnelControl(&route.Routes, web.IsFindRoute)
}
