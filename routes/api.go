package routes

import (
	"attrtour/core"
	"fmt"
	"net/http"
)

type Api struct{}

func (api *Api) Run(req *http.Request) {
	route := core.Route{
		Req: *req,
	}

	// User paths
	route.Get("/api/user/get", func() {
		fmt.Println("Get user by Id=578")
	})

	route.Post("/api/user/add", func() {
		fmt.Println("Add new user")
	})
}
