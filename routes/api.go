package routes

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"attrtour/app/controllers"
	"attrtour/core"
)

type Api struct {
	DBLink *sql.DB
	Ctx    context.Context
}

func (api *Api) Run(wri http.ResponseWriter, req *http.Request) {
	router := core.HandleRouter{
		Req: req,
		Wri: wri,
	}
	route := core.Route{
		Routes: map[string][]interface{}{},
	}

	// User paths
	route.Get("/api/user/get/{id}/", func(args map[string]interface{}) {
		uc := &controllers.UserController{
			Req: req,
			Wri: wri,
		}
		id, _ := strconv.Atoi(fmt.Sprintf("%v", args["id"]))

		uc.Get(id)
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
