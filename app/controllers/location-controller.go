package controllers

import (
	"encoding/json"
	"io"

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

	ls.Create(dtoData)
}

func (lc *LocationController) Save() {
	b, _ := io.ReadAll(lc.Req.Body)
	dtoData := dto.SaveLocationDto{}
	_ = json.Unmarshal(b, &dtoData)

	ls := services.LocationService{}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

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
