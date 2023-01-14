package controllers

import (
	"fmt"
	"io"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"
)

type UserController core.BaseController

func (uc *UserController) Add() {
	b, _ := io.ReadAll(uc.Req.Body)
	reqBody := string(b)
	fmt.Println(reqBody)
	fmt.Println()
}

func (uc *UserController) Get(id int) {
	um := models.UserModel{
		Dto: dto.GetUserDto{
			Id: id,
		},
	}
	um.RequestInit(uc.Wri, uc.Req)

	um.GetId()
}
