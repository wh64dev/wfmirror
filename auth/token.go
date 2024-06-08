package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/wh64dev/wfcloud/config"
)

type Claims struct {
	TokenID  string `json:"token"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (acc *Account) GenToken() (*string, error) {
	cnf := config.Get()
	claims := Claims{
		TokenID:  uuid.NewString(),
		UserID:   acc.Id,
		Username: acc.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	signed, err := token.SignedString([]byte(cnf.JWT.PrivKey))
	if err != nil {
		return nil, err
	}

	return &signed, nil
}

func Verifier(token string) (*Claims, error) {
	cnf := config.Get()
	claims := Claims{}
	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(cnf.JWT.PrivKey), nil
	}

	ref, err := jwt.ParseWithClaims(token, &claims, key)
	if err != nil {
		return nil, err
	}

	if !ref.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}
