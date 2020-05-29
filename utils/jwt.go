package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTSecret = []byte("API-SECRET~")

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["iat"] = time.Now().Unix()
	//claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["exp"] = time.Now().AddDate(1, 0, 0).Add(time.Second * 10).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
