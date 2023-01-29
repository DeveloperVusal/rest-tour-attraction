package services

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct{}

type Claims struct {
	jwt.RegisteredClaims
	UserId uint `json:"user_id"`
}

func (a *Auth) CreateToken(id uint, secret string) (string, error) {
	initToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			Issuer:    strconv.FormatInt(time.Now().Unix(), 10),
		},
	})
	var signedKey interface{} = []byte(secret)

	return initToken.SignedString(signedKey)
}
