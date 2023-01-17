package controllers

import (
	"encoding/json"
	"io"

	"attrtour/app/http/dto"
	"attrtour/app/services"
	"attrtour/core"
)

type GroupController core.BaseController

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

func (gc *GroupController) Get(id uint) {
	gs := services.GroupService{
		GetDto: dto.GetUserDto{
			Id: id,
		},
	}
	gs.RequestInit(gc.DBLink, gc.Wri, gc.Req)

	gs.GetById()
}
