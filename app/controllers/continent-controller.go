package controllers

import (
	"encoding/json"
	"io"

	"attrtour/app/http/dto"
	"attrtour/app/services"
	"attrtour/core"
)

type ContinentController core.BaseController

func (cc *ContinentController) Get() {
	cs := services.ContinentService{}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)
	cs.GetAll()
}

func (cc *ContinentController) GetById(id uint) {
	cs := services.ContinentService{
		GetDto: dto.GetContinentDto{
			Id: id,
		},
	}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)

	cs.GetById()
}

func (cc *ContinentController) Add() {
	b, _ := io.ReadAll(cc.Req.Body)
	dtoData := dto.AddContinentDto{}
	_ = json.Unmarshal(b, &dtoData)

	cs := services.ContinentService{}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)

	cs.Create(dtoData)
}

func (cc *ContinentController) Save() {
	b, _ := io.ReadAll(cc.Req.Body)
	dtoData := dto.SaveContinentDto{}
	_ = json.Unmarshal(b, &dtoData)

	cs := services.ContinentService{}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)

	cs.Save(dtoData)
}

func (cc *ContinentController) Delete(id uint) {
	cs := services.ContinentService{
		GetDto: dto.GetContinentDto{
			Id: id,
		},
	}
	cs.RequestInit(cc.DBLink, cc.Wri, cc.Req)

	cs.Delete()
}
