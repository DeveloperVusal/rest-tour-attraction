package controllers

import (
	"encoding/json"
	"io"

	"attrtour/app/http/dto"
	"attrtour/app/services"
	"attrtour/core"
)

type GroupController core.BaseController

func (gc *GroupController) Get() {
	gs := services.GroupService{}
	gs.RequestInit(gc.DBLink, gc.Wri, gc.Req)

	gs.GetAll()
}

func (gc *GroupController) GetById(id uint) {
	gs := services.GroupService{
		GetDto: dto.GetGroupDto{
			Id: id,
		},
	}
	gs.RequestInit(gc.DBLink, gc.Wri, gc.Req)

	gs.GetById()
}

func (gc *GroupController) Add() {
	b, _ := io.ReadAll(gc.Req.Body)
	dtoData := dto.AddGroupDto{}
	_ = json.Unmarshal(b, &dtoData)

	gs := services.GroupService{}
	gs.RequestInit(gc.DBLink, gc.Wri, gc.Req)

	gs.Create(dtoData)
}

func (gc *GroupController) Save() {
	b, _ := io.ReadAll(gc.Req.Body)
	dtoData := dto.SaveGroupDto{}
	_ = json.Unmarshal(b, &dtoData)

	gs := services.GroupService{}
	gs.RequestInit(gc.DBLink, gc.Wri, gc.Req)

	gs.Save(dtoData)
}

func (gc *GroupController) Delete(id uint) {
	gs := services.GroupService{
		GetDto: dto.GetGroupDto{
			Id: id,
		},
	}
	gs.RequestInit(gc.DBLink, gc.Wri, gc.Req)

	gs.Delete()
}
