package controllers

import (
	"encoding/json"
	"io"

	"attrtour/app/http/dto"
	"attrtour/app/services"
	"attrtour/core"
)

type UserController core.BaseController

func (uc *UserController) Get() {
	us := services.UserService{}
	us.RequestInit(uc.DBLink, uc.Wri, uc.Req)

	us.GetAll()
}

func (uc *UserController) GetById(id uint) {
	us := services.UserService{
		GetDto: dto.GetUserDto{
			Id: id,
		},
	}
	us.RequestInit(uc.DBLink, uc.Wri, uc.Req)

	us.GetById()
}

func (uc *UserController) Add() {
	b, _ := io.ReadAll(uc.Req.Body)
	dtoData := dto.AddUserDto{}
	_ = json.Unmarshal(b, &dtoData)

	us := services.UserService{}
	us.RequestInit(uc.DBLink, uc.Wri, uc.Req)

	us.Create(dtoData)
}

func (uc *UserController) Save() {
	b, _ := io.ReadAll(uc.Req.Body)
	dtoData := dto.SaveUserDto{}
	_ = json.Unmarshal(b, &dtoData)

	us := services.UserService{}
	us.RequestInit(uc.DBLink, uc.Wri, uc.Req)

	us.Save(dtoData)
}

func (uc *UserController) Delete(id uint) {
	us := services.UserService{
		GetDto: dto.GetUserDto{
			Id: id,
		},
	}
	us.RequestInit(uc.DBLink, uc.Wri, uc.Req)

	us.Delete()
}
