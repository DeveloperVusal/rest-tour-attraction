package services

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"
	// "attrtour/app/models"
)

type AuthService struct {
	core.BaseService
}

type Claims struct {
	jwt.RegisteredClaims
	UserId uint `json:"user_id"`
}

func (as *AuthService) Login(_dto dto.AuthDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		var user models.User

		env := core.Helpers{}
		us := UserService{}

		if env.Env("APP_ROOT_LOGIN") == _dto.Username {
			hashPasswd, _ := us.HashPassword(env.Env("APP_ROOT_PASSWD"))

			if us.CheckPasswordHash(_dto.Password, hashPasswd) {
				as.DBLink.First(&user, "username = ?", _dto.Username)

				access, refresh := as.CreatePairTokens(user.Id, env.Env("APP_JWT_SECRET"))
				auth := models.Auth{
					UserId:       user.Id,
					AccessToken:  access,
					RefreshToken: refresh,
					Ip:           as.Req.RemoteAddr,
					Device:       "site",
					UserAgent:    as.Req.UserAgent(),
				}

				as.DBLink.Save(&auth)

				jsonData, _ := json.Marshal(map[string]interface{}{
					"access_token":  access,
					"refresh_token": refresh,
				})

				as.Wri.Write(jsonData)
			} else {
				jsonData, _ := json.Marshal(map[string]interface{}{
					"status":  "warning",
					"message": "incorrect data",
				})

				as.Wri.Write(jsonData)
			}
		} else {
			as.DBLink.Model(models.User{Username: _dto.Username}).First(&user)

			if us.CheckPasswordHash(_dto.Password, user.Password) {
				access, refresh := as.CreatePairTokens(user.Id, env.Env("APP_JWT_SECRET"))

				auth := models.Auth{
					UserId:       user.Id,
					AccessToken:  access,
					RefreshToken: refresh,
					Ip:           as.Req.RemoteAddr,
					Device:       "site",
					UserAgent:    as.Req.UserAgent(),
				}

				as.DBLink.Save(&auth)

				jsonData, _ := json.Marshal(map[string]interface{}{
					"access_token":  access,
					"refresh_token": refresh,
				})

				as.Wri.Write(jsonData)
			} else {
				jsonData, _ := json.Marshal(map[string]interface{}{
					"status":  "warning",
					"message": "incorrect data",
				})

				as.Wri.Write(jsonData)
			}
		}
	} else {
		as.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (as *AuthService) CreatePairTokens(id uint, secret string) (string, string) {
	var signedKey interface{} = []byte(secret)

	initToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			Issuer:    strconv.FormatInt(time.Now().Unix(), 10),
		},
	})

	access, _ := initToken.SignedString(signedKey)

	initToken = jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 43830)),
			Issuer:    strconv.FormatInt(time.Now().Unix(), 10),
		},
	})
	refresh, _ := initToken.SignedString(signedKey)

	return access, refresh
}
