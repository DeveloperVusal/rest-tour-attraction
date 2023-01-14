package controllers

import (
	"encoding/json"
	"io"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"
)

type UserController core.BaseController

func (uc *UserController) Add() {
	b, _ := io.ReadAll(uc.Req.Body)
	dtoData := dto.AddUserDto{}
	_ = json.Unmarshal(b, &dtoData)

	um := models.UserModel{}
	um.RequestInit(uc.DBLink, uc.Ctx, uc.Wri, uc.Req)

	um.Create(dtoData)
}

func (uc *UserController) Save() {
	b, _ := io.ReadAll(uc.Req.Body)
	dtoData := dto.SaveUserDto{}
	_ = json.Unmarshal(b, &dtoData)

	um := models.UserModel{}
	um.RequestInit(uc.DBLink, uc.Ctx, uc.Wri, uc.Req)

	um.Save(dtoData)
}

func (uc *UserController) Get(id int) {
	um := models.UserModel{
		GetDto: dto.GetUserDto{
			Id: id,
		},
	}
	um.RequestInit(uc.DBLink, uc.Ctx, uc.Wri, uc.Req)

	um.GetId()
}
