package routes

import (
	"attrtour/core"
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

type Api struct {
	DBLink *sql.DB
	Ctx    context.Context
}

func (api *Api) Run(wri http.ResponseWriter, req *http.Request) {
	router := core.HandleRouter{
		Req: *req,
		Wri: wri,
	}
	route := core.Route{
		Routes: map[string][]interface{}{},
	}

	// User paths
	route.Get("/api/user/get/{section}/{id}/", func(args map[string]interface{}) {
		wri.Write([]byte("Get user section=" + fmt.Sprintf("%v", args["section"]) + "\n"))
		wri.Write([]byte("Get user by Id=" + fmt.Sprintf("%v", args["id"]) + "\n"))
	})

	route.Post("/api/user/add", func(args map[string]interface{}) {
		wri.Write([]byte("Add new user"))
	})

	route.Post("/api/place/add", func(args map[string]interface{}) {
		wri.Write([]byte("Add new place"))
	})

	// Run Router Handle
	router.TunnelControl(&route.Routes)
}
