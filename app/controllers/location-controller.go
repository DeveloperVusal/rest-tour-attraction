package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"attrtour/app/http/dto"
	"attrtour/app/services"
	"attrtour/core"
)

type LocationController core.BaseController

func (lc *LocationController) Get() {
	ls := services.LocationService{}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	ls.GetAll()
}

func (lc *LocationController) GetById(id uint) {
	ls := services.LocationService{
		GetDto: dto.GetLocationDto{
			Id: id,
		},
	}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	ls.GetById()
}

func (lc *LocationController) Add() {
	b, _ := io.ReadAll(lc.Req.Body)
	dtoData := dto.AddLocationDto{}
	_ = json.Unmarshal(b, &dtoData)

	ls := services.LocationService{}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	as := services.AuthService{}
	authToken := lc.Req.Header["Authorization"]
	split := strings.Split(authToken[0], " ")
	claims, ok := as.GetClaims(split[1])

	if ok {
		convUint, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
		dtoData.UserId = uint(convUint)
	}

	ls.Create(dtoData)
}

func (lc *LocationController) Save() {
	b, _ := io.ReadAll(lc.Req.Body)
	dtoData := dto.SaveLocationDto{}
	_ = json.Unmarshal(b, &dtoData)

	ls := services.LocationService{}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	as := services.AuthService{}
	authToken := lc.Req.Header["Authorization"]
	split := strings.Split(authToken[0], " ")
	claims, ok := as.GetClaims(split[1])

	if ok {
		convUint, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
		dtoData.UserId = uint(convUint)
	}

	ls.Save(dtoData)
}

func (lc *LocationController) Delete(id uint) {
	ls := services.LocationService{
		GetDto: dto.GetLocationDto{
			Id: id,
		},
	}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	ls.Delete()
}
