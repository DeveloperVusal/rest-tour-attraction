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

	// Group paths
	route.Get("/api/group/get", func(args map[string]interface{}) {
		gc := &controllers.GroupController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Get()
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
			gc.GetById(uint(id))
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

	route.Delete("/api/group/delete/{id}/", func(args map[string]interface{}) {
		gc := &controllers.GroupController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			gc.Delete(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	// User paths
	route.Get("/api/user/get", func(args map[string]interface{}) {
		uc := &controllers.UserController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		uc.Get()
	})

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
			uc.GetById(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	route.Delete("/api/user/delete/{id}/", func(args map[string]interface{}) {
		uc := &controllers.UserController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			uc.Delete(uint(id))
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

	// Language paths
	route.Get("/api/language/get", func(args map[string]interface{}) {
		lc := &controllers.LanguageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		lc.Get()
	})

	route.Get("/api/language/get/{id}/", func(args map[string]interface{}) {
		lc := &controllers.LanguageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			lc.GetById(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	route.Post("/api/language/add", func(args map[string]interface{}) {
		lc := &controllers.LanguageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		lc.Add()
	})

	route.Put("/api/language/save", func(args map[string]interface{}) {
		lc := &controllers.LanguageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		lc.Save()
	})

	route.Delete("/api/language/delete/{id}/", func(args map[string]interface{}) {
		lc := &controllers.LanguageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			lc.Delete(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	// Continent paths
	route.Get("/api/continent/get", func(args map[string]interface{}) {
		gc := &controllers.ContinentController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Get()
	})

	route.Get("/api/continent/get/{id}/", func(args map[string]interface{}) {
		gc := &controllers.ContinentController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			gc.GetById(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	route.Post("/api/continent/add", func(args map[string]interface{}) {
		gc := &controllers.ContinentController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Add()
	})

	route.Put("/api/continent/save", func(args map[string]interface{}) {
		gc := &controllers.ContinentController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Save()
	})

	route.Delete("/api/continent/delete/{id}/", func(args map[string]interface{}) {
		gc := &controllers.ContinentController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			gc.Delete(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	// Country paths
	route.Get("/api/country/get", func(args map[string]interface{}) {
		cc := &controllers.CountryController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		cc.Get()
	})

	route.Get("/api/country/get/{id}/", func(args map[string]interface{}) {
		cc := &controllers.CountryController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			cc.GetById(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	route.Post("/api/country/add", func(args map[string]interface{}) {
		cc := &controllers.CountryController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		cc.Add()
	})

	route.Put("/api/country/save", func(args map[string]interface{}) {
		cc := &controllers.CountryController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		cc.Save()
	})

	route.Delete("/api/country/delete/{id}/", func(args map[string]interface{}) {
		cc := &controllers.CountryController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			cc.Delete(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	/*==========================*/

	// Run Router Handle
	router.TunnelControl(&route.Routes)
}
