package util

import (
	"errors"
	"git.woa.com/test/helloworld/internal/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtPayLoad struct {
	Appid  string `json:"appid"`
	Appkey string `json:"appkey"`
}

const secretKey = "dddddd"

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

func GenToken(user JwtPayLoad) (string, error) {
	var secretKey = []byte(secretKey)
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Second).Unix(), // 过期时间
			Issuer:    "wenli",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	var secretKey = []byte(secretKey)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func Auth(tokenStr string) error {
	user, err := ParseToken(tokenStr)
	if err != nil {
		return err
	}
	exist := models.CheckAuth(user.JwtPayLoad.Appid, user.JwtPayLoad.Appkey)
	if exist {
		return nil
	}
	return errors.New(user.JwtPayLoad.Appid + "查不到对应的appid" + user.JwtPayLoad.Appkey)
}
