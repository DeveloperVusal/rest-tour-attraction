package controllers

import (
	"encoding/json"
	"io"

	"attrtour/app/http/dto"
	"attrtour/app/services"
	"attrtour/core"
)

type CountryController core.BaseController

func (cc *CountryController) Get() {
	cs := services.CountryService{}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)

	cs.GetAll()
}

func (cc *CountryController) GetById(id uint) {
	cs := services.CountryService{
		GetDto: dto.GetCountryDto{
			Id: id,
		},
	}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)

	cs.GetById()
}

func (cc *CountryController) Add() {
	b, _ := io.ReadAll(cc.Req.Body)
	dtoData := dto.AddCountryDto{}
	_ = json.Unmarshal(b, &dtoData)

	cs := services.CountryService{}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)

	cs.Create(dtoData)
}

func (cc *CountryController) Save() {
	b, _ := io.ReadAll(cc.Req.Body)
	dtoData := dto.SaveCountryDto{}
	_ = json.Unmarshal(b, &dtoData)

	cs := services.CountryService{}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)

	cs.Save(dtoData)
}

func (cc *CountryController) Delete(id uint) {
	cs := services.CountryService{
		GetDto: dto.GetCountryDto{
			Id: id,
		},
	}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)

	cs.Delete()
}
