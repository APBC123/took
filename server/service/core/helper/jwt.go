package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Id       int64
	Username string
	Password string
	jwt.StandardClaims
}

func GenerateToken(id int64, username, password, secretKey string, second int64) (string, error) {
	uc := UserClaim{
		Id:       id,
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + second,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyzeToken(token, secretKey string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if !claims.Valid {
		return uc, errors.New("用户Token不合法")
	}
	if err != nil {
		return nil, err
	}
	return uc, err
}