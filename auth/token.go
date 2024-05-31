package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wh64dev/wfcloud/config"
)

type Claims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var exp = 1 * time.Hour

func (acc *Account) GenToken() (string, error) {
	cnf := config.Get()
	exp := time.Now().Add(exp)
	claims := &Claims{
		Id:       acc.Id,
		Username: acc.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cnf.JWT.JWTToken))
}
