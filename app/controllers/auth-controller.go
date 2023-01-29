package controllers

import (
	"attrtour/app/services"
	"attrtour/core"
)

type AuthController core.BaseController

func (ac *AuthController) Login() {
	as := services.Auth{}
	token, _ := as.CreateToken(2, "gTphmKeR2fAbNTasWY90pDwsXOEFZWTT")

	ac.Wri.Write([]byte(token))

}
func (ac *AuthController) VerifyToken() {

}

func (ac *AuthController) Refresh() {

}

func (ac *AuthController) Logout() {

}
