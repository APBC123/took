package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey = "took"

type UserClaim struct {
	Id       int
	Username string
	Password string
	jwt.StandardClaims
}

func GenerateToken(id int, username string, password string, second int) (string, error) {
	uc := UserClaim{
		Id:       id,
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyzeToken(token string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}
