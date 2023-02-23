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
func (ac *AuthController) Logout() {
	as := services.AuthService{}
	as.RequestInit(ac.DBLink, ac.Wri, ac.Req)

	as.Logout(ac.Req.Header["Authorization"])
}
func (ac *AuthController) VerifyToken() {
	as := services.AuthService{}
	as.RequestInit(ac.DBLink, ac.Wri, ac.Req)

	as.VerifyToken(ac.Req.Header["Authorization"])
}

func (ac *AuthController) Refresh() {
	token := ac.Req.FormValue("token")

	if len(token) > 50 {
		as := services.AuthService{}
		as.RequestInit(ac.DBLink, ac.Wri, ac.Req)

		as.Refresh(token)
	}
}
