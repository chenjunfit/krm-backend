package jwtutil

import (
	"github.com/golang-jwt/jwt/v5"
	"krm-backend/config"
	"krm-backend/utils/logs"
	"time"
)

var jwtSignKey = []byte(config.JwtSignKey)

type MyCustomClaims struct {
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

func GenToken(userName string) (string, error) {
	claims := MyCustomClaims{
		userName,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.JwtExpireTime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "server",
			Subject:   "server-token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(jwtSignKey)
	return ss, err
}

func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSignKey), nil
	})
	if err != nil {
		logs.Error(nil, "token解析失败")
		return nil, err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims, nil
	} else {
		logs.Error(nil, "token不合法")
		return nil, err
	}
}
