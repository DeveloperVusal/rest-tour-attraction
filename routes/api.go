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
	DBLink      *gorm.DB
	IsFindRoute *bool
}

func (api *Api) Run(wri http.ResponseWriter, req *http.Request) {
	router := core.HandleRouter{
		Req: req,
		Wri: wri,
	}
	route := core.Route{
		Routes: map[string][]interface{}{},
		Prefix: "/api",
	}

	// Group paths
	route.Get("/group/get", func(args map[string]interface{}) {
		gc := &controllers.GroupController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Get()
	})

	route.Get("/group/get/{id}/", func(args map[string]interface{}) {
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

	route.Post("/group/add", func(args map[string]interface{}) {
		gc := &controllers.GroupController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Add()
	})

	route.Patch("/group/save", func(args map[string]interface{}) {
		gc := &controllers.GroupController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Save()
	})

	route.Delete("/group/delete/{id}/", func(args map[string]interface{}) {
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
	route.Get("/user/get", func(args map[string]interface{}) {
		uc := &controllers.UserController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		uc.Get()
	})

	route.Get("/user/get/{id}/", func(args map[string]interface{}) {
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

	route.Delete("/user/delete/{id}/", func(args map[string]interface{}) {
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

	route.Post("/user/add", func(args map[string]interface{}) {
		uc := &controllers.UserController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		uc.Add()
	})

	route.Patch("/user/save", func(args map[string]interface{}) {
		uc := &controllers.UserController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		uc.Save()
	})

	// Language paths
	route.Get("/language/get", func(args map[string]interface{}) {
		lc := &controllers.LanguageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		lc.Get()
	})

	route.Get("/language/get/{id}/", func(args map[string]interface{}) {
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

	route.Post("/language/add", func(args map[string]interface{}) {
		lc := &controllers.LanguageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		lc.Add()
	})

	route.Patch("/language/save", func(args map[string]interface{}) {
		lc := &controllers.LanguageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		lc.Save()
	})

	route.Delete("/language/delete/{id}/", func(args map[string]interface{}) {
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
	route.Get("/continent/get", func(args map[string]interface{}) {
		gc := &controllers.ContinentController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Get()
	})

	route.Get("/continent/get/{id}/", func(args map[string]interface{}) {
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

	route.Post("/continent/add", func(args map[string]interface{}) {
		gc := &controllers.ContinentController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Add()
	})

	route.Patch("/continent/save", func(args map[string]interface{}) {
		gc := &controllers.ContinentController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		gc.Save()
	})

	route.Delete("/continent/delete/{id}/", func(args map[string]interface{}) {
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
	route.Get("/country/get", func(args map[string]interface{}) {
		cc := &controllers.CountryController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		cc.Get()
	})

	route.Get("/country/get/{id}/", func(args map[string]interface{}) {
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

	route.Post("/country/add", func(args map[string]interface{}) {
		cc := &controllers.CountryController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		cc.Add()
	})

	route.Patch("/country/save", func(args map[string]interface{}) {
		cc := &controllers.CountryController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		cc.Save()
	})

	route.Delete("/country/delete/{id}/", func(args map[string]interface{}) {
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

	// Location paths
	route.Get("/location/get", func(args map[string]interface{}) {
		lc := &controllers.LocationController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		lc.Get()
	})

	route.Get("/location/get/{id}/", func(args map[string]interface{}) {
		lc := &controllers.LocationController{
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

	route.Post("/location/add", func(args map[string]interface{}) {
		lc := &controllers.LocationController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		lc.Add()
	})

	route.Patch("/location/save", func(args map[string]interface{}) {
		lc := &controllers.LocationController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		lc.Save()
	})

	route.Delete("/location/delete/{id}/", func(args map[string]interface{}) {
		lc := &controllers.LocationController{
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

	// Upload paths
	route.Post("/main-image/upload", func(args map[string]interface{}) {
		mic := &controllers.MainImageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		mic.Upload()
	})

	route.Patch("/main-image/save", func(args map[string]interface{}) {
		mic := &controllers.MainImageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		mic.Save()
	})

	route.Delete("/main-image/remove/{id}/", func(args map[string]interface{}) {
		mic := &controllers.MainImageController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		idStr := fmt.Sprintf("%v", args["id"])
		match, _ := regexp.MatchString("[0-9]", idStr)

		if match {
			id, _ := strconv.Atoi(idStr)
			mic.Remove(uint(id))
		} else {
			wri.Write([]byte(fmt.Sprintf("Invalid object type: expected `int`, turned out to be `%T`", args["id"])))
		}
	})

	/*==========================*/

	//Auth
	route.Post("/account/auth", func(args map[string]interface{}) {
		ac := &controllers.AuthController{
			Req:    req,
			Wri:    wri,
			DBLink: api.DBLink,
		}

		ac.Login()
	})

	// Run Router Handle
	router.TunnelControl(&route.Routes, api.IsFindRoute)
}
