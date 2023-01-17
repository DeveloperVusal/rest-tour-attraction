package routes

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"attrtour/app/controllers"
	"attrtour/core"

	"gorm.io/gorm"
)

type Api struct {
	DBLink *gorm.DB
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
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			uc.Get(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	route.Post("/api/user/add", func(args map[string]interface{}) {
		uc := &controllers.UserController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		uc.Add()
	})

	route.Put("/api/user/save", func(args map[string]interface{}) {
		uc := &controllers.UserController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		uc.Save()
	})

	route.Get("/api/group/get/{id}/", func(args map[string]interface{}) {
		gc := &controllers.GroupController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			gc.Get(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	route.Post("/api/group/add", func(args map[string]interface{}) {
		gc := &controllers.GroupController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Add()
	})

	route.Put("/api/group/save", func(args map[string]interface{}) {
		gc := &controllers.GroupController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Save()
	})

	// Run Router Handle
	router.TunnelControl(&route.Routes)
}
