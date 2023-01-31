package controllers

import (
	"attrtour/app/http/dto"
	"attrtour/app/services"
	"attrtour/core"
	"encoding/json"
	"io"
)

type AuthController core.BaseController

func (ac *AuthController) Login() {
	b, _ := io.ReadAll(ac.Req.Body)
	dtoData := dto.AuthDto{}
	_ = json.Unmarshal(b, &dtoData)

	as := services.AuthService{}
	as.RequestInit(ac.DBLink, ac.Wri, ac.Req)

	as.Login(dtoData)

}
func (ac *AuthController) VerifyToken() {

}

func (ac *AuthController) Refresh() {

}

func (ac *AuthController) Logout() {

}
