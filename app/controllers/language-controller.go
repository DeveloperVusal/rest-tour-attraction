package controllers

import (
	"encoding/json"
	"io"

	"attrtour/app/http/dto"
	"attrtour/app/services"
	"attrtour/core"
)

type LanguageController core.BaseController

func (lc *LanguageController) Get() {
	ls := services.LanguageService{}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	ls.GetAll()
}

func (lc *LanguageController) GetById(id uint) {
	ls := services.LanguageService{
		GetDto: dto.GetLanguageDto{
			Id: id,
		},
	}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	ls.GetById()
}

func (lc *LanguageController) Add() {
	b, _ := io.ReadAll(lc.Req.Body)
	dtoData := dto.AddLanguageDto{}
	_ = json.Unmarshal(b, &dtoData)

	ls := services.LanguageService{}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	ls.Create(dtoData)
}

func (lc *LanguageController) Save() {
	b, _ := io.ReadAll(lc.Req.Body)
	dtoData := dto.SaveLanguageDto{}
	_ = json.Unmarshal(b, &dtoData)

	ls := services.LanguageService{}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	ls.Save(dtoData)
}

func (lc *LanguageController) Delete(id uint) {
	ls := services.LanguageService{
		GetDto: dto.GetLanguageDto{
			Id: id,
		},
	}
	ls.RequestInit(lc.DBLink, lc.Wri, lc.Req)

	ls.Delete()
}
