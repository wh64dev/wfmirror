package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Id       string
	Username string
	jwt.StandardClaims
}

var (
	exp     = 1 * time.Hour
	PrivKey = []byte(os.Getenv("TOKEN"))
)

func (acc *Account) GenToken() (string, error) {
	exp := time.Now().Add(exp)
	claims := &Claims{
		Id:       acc.GetId(),
		Username: acc.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(PrivKey)
}
