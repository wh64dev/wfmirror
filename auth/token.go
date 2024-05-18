package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
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
		Username: acc.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(PrivKey)
}
