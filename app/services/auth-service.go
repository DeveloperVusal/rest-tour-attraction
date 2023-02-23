package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"
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

func (as *AuthService) Logout(_auth []string) {
	var auth models.Auth

	split := strings.Split(_auth[0], " ")
	tokenString := split[1]

	as.DBLink.Where("access_token = ?", tokenString).Delete(&auth)

	as.Wri.WriteHeader(http.StatusOK)
}

func (as *AuthService) VerifyToken(_auth []string) {
	var auth models.Auth

	split := strings.Split(_auth[0], " ")
	tokenString := split[1]
	isVerify, _ := as.IsVerifyToken(tokenString)

	as.DBLink.First(&auth, "access_token = ?", tokenString)

	if auth.Id <= 0 {
		isVerify = false
	}

	if isVerify {
		as.Wri.WriteHeader(http.StatusOK)
	} else {
		as.Wri.WriteHeader(http.StatusUnauthorized)
	}
}

func (as *AuthService) IsVerifyToken(tokenString string) (bool, *jwt.Token) {
	env := core.Helpers{}

	var signedKey interface{} = []byte(env.Env("APP_JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return signedKey, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, token
	} else {
		fmt.Println(err)
		return false, token
	}
}

func (as *AuthService) GetClaims(tokenString string) (jwt.MapClaims, bool) {
	env := core.Helpers{}

	var signedKey interface{} = []byte(env.Env("APP_JWT_SECRET"))

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return signedKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	return claims, ok
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

func (as *AuthService) Refresh(tokenString string) {
	env := core.Helpers{}
	isVerify, jwtToken := as.IsVerifyToken(tokenString)

	if isVerify {
		var auth models.Auth
		var jsonData []byte

		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		userId, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
		access, refresh := as.CreatePairTokens(uint(userId), env.Env("APP_JWT_SECRET"))

		as.DBLink.First(&auth, "refresh_token = ?", tokenString)

		if auth.Id > 0 {
			auth.RefreshToken = refresh
			auth.AccessToken = access

			as.DBLink.Save(&auth)

			jsonData, _ = json.Marshal(map[string]interface{}{
				"access_token":  access,
				"refresh_token": refresh,
			})
		} else {
			jsonData, _ = json.Marshal(map[string]interface{}{
				"status":  "error",
				"message": "Invalid token",
			})
		}
		as.Wri.Write(jsonData)

	} else {
		jsonData, _ := json.Marshal(map[string]interface{}{
			"status":  "error",
			"message": "Token is expired",
		})

		as.Wri.Write(jsonData)
	}
}
